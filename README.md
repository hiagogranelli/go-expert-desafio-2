# Desafio Consulta CEP Concorrente

Este projeto em Go implementa uma solução para consultar informações de um CEP (Código de Endereçamento Postal) em duas APIs distintas simultaneamente, retornando o resultado daquela que responder mais rápido.

## Funcionalidades

*   **Consulta Concorrente:** Realiza requisições HTTP GET para as APIs [BrasilAPI](https://brasilapi.com.br/) e [ViaCEP](http://viacep.com.br/) ao mesmo tempo.
*   **Resposta Mais Rápida:** Utiliza goroutines e channels para determinar qual API respondeu primeiro.
*   **Descarte da Lenta:** A resposta da API mais lenta é descartada.
*   **Timeout:** Implementa um limite de tempo de 1 segundo para obter uma resposta. Se nenhuma API responder dentro desse tempo, uma mensagem de erro é exibida.
*   **Exibição no Console:** O resultado da API vencedora (contendo os dados do endereço) e o nome da API são impressos no terminal.

## APIs Utilizadas

*   BrasilAPI: `https://brasilapi.com.br/api/cep/v1/{cep}`
*   ViaCEP: `http://viacep.com.br/ws/{cep}/json/`

## Pré-requisitos

*   Go (versão 1.18 ou superior recomendado) instalado em sua máquina. ([Instruções de Instalação](https://go.dev/doc/install))

## Como Usar

1.  **Clone o repositório ou salve o código:**
    *   Se clonou: `git clone <URL_DO_REPOSITORIO>` e entre na pasta `cd <NOME_DA_PASTA>`.
    *   Se salvou: Certifique-se de que o arquivo `main.go` está no seu diretório atual.

2.  **Execute o programa:**
    Abra seu terminal ou prompt de comando e execute:

    ```bash
    go run main.go SEU_CEP_AQUI
    ```

    *   **Substitua `SEU_CEP_AQUI`** pelo CEP que deseja consultar (somente números, por exemplo: `01153000`).

## Exemplo de Saída

```
Buscando CEP: 06230100
Iniciando busca nas APIs...

--- Resultado (API mais rápida: ViaCEP) ---
{
  "cep": "06230-100",
  "logradouro": "Rua Amador Bueno",
  "complemento": "",
  "unidade": "",
  "bairro": "Piratininga",
  "localidade": "Osasco",
  "uf": "SP",
  "estado": "São Paulo",
  "regiao": "Sudeste",
  "ibge": "3534401",
  "gia": "4923",
  "ddd": "11",
  "siafi": "6789"
}

Busca finalizada.
```

