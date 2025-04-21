package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Result struct {
	SourceAPI string
	Response  string
	Error     error
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Erro: Forneça o CEP como argumento.")
		fmt.Fprintln(os.Stderr, "Uso: go run main.go <cep>")
		os.Exit(1)
	}
	cep := os.Args[1]

	if len(cep) < 8 {
		fmt.Fprintln(os.Stderr, "Erro: CEP inválido.")
		os.Exit(1)
	}

	brasilAPIURL := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
	viaCEPURL := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)

	resultChan := make(chan Result)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	fmt.Println("Buscando CEP:", cep)
	fmt.Println("Iniciando busca nas APIs...")

	go fetchAPI(ctx, "BrasilAPI", brasilAPIURL, resultChan)
	go fetchAPI(ctx, "ViaCEP", viaCEPURL, resultChan)

	select {
	case res := <-resultChan:
		if res.Error != nil {
			fmt.Fprintf(os.Stderr, "Erro recebido da API %s: %v\n", res.SourceAPI, res.Error)
			os.Exit(1)
		}
		fmt.Printf("\n--- Resultado (API mais rápida: %s) ---\n", res.SourceAPI)
		fmt.Println(res.Response)

	case <-ctx.Done():
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Fprintln(os.Stderr, "\nErro: Timeout de 1 segundo atingido. Nenhuma API respondeu a tempo.")
			os.Exit(1)
		} else {
			fmt.Fprintf(os.Stderr, "\nErro: Contexto cancelado: %v\n", ctx.Err())
			os.Exit(1)
		}
	}

	fmt.Println("\nBusca finalizada.")
}

func fetchAPI(ctx context.Context, apiName string, url string, resultChan chan<- Result) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	result := Result{
		SourceAPI: apiName,
		Response:  string(bodyBytes),
		Error:     nil,
	}

	select {
	case resultChan <- result:
	case <-ctx.Done():
	}
}
