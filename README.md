# Password Generator

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

## Descrição

O **Password Generator** é uma ferramenta de linha de comando escrita em Go que gera senhas aleatórias com base em critérios especificados pelo usuário, como comprimento e tipos de caracteres (maiúsculas, minúsculas, números e símbolos).

## Funcionalidades

- Geração de senhas seguras e aleatórias.
- Opções para incluir/excluir maiúsculas, minúsculas, números e símbolos.
- Validação das senhas geradas para garantir que atendem aos critérios especificados.

## Instalação

1. Clone o repositório:
    ```sh
    git clone https://github.com/fabiosoliveira/password-generator.git
    cd password-generator
    ```

2. Compile o projeto:
    ```sh
    go build -o gen-pass cmd/cli/main.go
    ```

## Uso

Execute o comando `gen-pass` seguido dos parâmetros:

```sh
./gen-pass <length:number> <uppercase:bool> <lowercase:bool> <numbers:bool> <symbols:bool>
```

## Exemplo
Para gerar uma senha de 16 caracteres contendo maiúsculas, minúsculas, números e símbolos:

```sh
./gen-pass 16 true true true true
```

## Estrutura do Projeto

```sh
password-generator/
├── .gitignore
├── cmd/
│   └── cli/
│       └── main.go
├── go.mod
├── LICENSE
├── pkg/
│   └── password/
│       ├── password.go
│       ├── password_test.go
│       ├── validator.go
│       └── validator_test.go
└── README.md
```


## Testes

Para executar os testes, use o comando:

```sh
go test ./pkg/password/...
```

## Licença
Este projeto está licenciado sob a licença MIT. Veja o arquivo LICENSE para mais detalhes.

## Contribuição
Contribuições são bem-vindas! Sinta-se à vontade para abrir issues e pull requests.

## Autor
Fabio S. Oliveira
Feito com ❤️ por Fabio S. Oliveira