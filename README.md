# Conversor de Temperatura API

Esta é uma API simples em Go para converter temperaturas entre Celsius, Fahrenheit e Kelvin.

## Requisitos

- Go (versão 1.13 ou superior)

## Instalação

1. Clone o repositório:

    ```bash
    git clone https://github.com/marcuscarvalhodev/conversorTemperaturaAPI.git
    cd conversorTemperaturaAPI
    ```

2. Instale as dependências:

    ```bash
    go mod tidy
    ```

3. Execute a aplicação:

    ```bash
    go run main.go
    ```

A API estará disponível em [http://localhost:8080](http://localhost:8080).

## Endpoints

### `POST /converter`

Realiza a conversão de temperatura.

#### Parâmetros

- `unidade_original`: Unidade de temperatura original (celsius, fahrenheit, kelvin).
- `unidade_convertida`: Unidade para a qual converter (celsius, fahrenheit, kelvin).
- `temperatura_inserida`: Temperatura a ser convertida.

#### Exemplo de Uso

```bash
curl -X POST -d "unidade_original=celsius&unidade_convertida=fahrenheit&temperatura_inserida=25" http://localhost:8080/converter
```

## Contribuição

Sinta-se à vontade para contribuir! Crie um fork do projeto, faça as alterações desejadas e envie um pull request.

