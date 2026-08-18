# Como executar o projeto

## Requisitos

- Go 1.26.1 ou versão compatível com o `go.mod`.
- Python 3 disponível pelo comando `python3`.

## Execução

Na raiz do projeto, execute:

```bash
go run .
```

O menu oferece os exemplos de WaitGroup, Channels, monitoramento de uptime e
interoperabilidade entre Go e Python. Na opção `5`, digite um texto; o programa
Go o enviará como argumento para `internal/interop/script.py` e imprimirá a
versão convertida para maiúsculas pelo Python.

A opção `4` executa todos os exemplos e também solicita o texto usado na
demonstração de interoperabilidade.

## Testes

```bash
go test ./...
go vet ./...
```
