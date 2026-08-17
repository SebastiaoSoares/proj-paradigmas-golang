# Instalação e execução

Este guia descreve como reproduzir os exemplos do repositório a partir de sua raiz.

## Pré-requisitos

- Go na versão declarada em [`go.mod`](go.mod);
- Python 3.9 ou posterior para a implementação com `threading` e `queue`;
- para o comparativo em C: sistema compatível com POSIX e compilador C com suporte a pthreads (`cc`, GCC ou Clang);
- para a verificação automática: shell POSIX e `diff`.

Confirme as ferramentas disponíveis:

```bash
go version
python3 --version
cc --version
```

O módulo usa apenas a biblioteca padrão de Go, portanto não há dependências Go de terceiros para instalar.

## Estudo de caso: monitoramento de URLs

```bash
go run .
```

O programa faz requisições de rede concorrentes. É necessário acesso à internet; códigos HTTP e mensagens de erro podem variar conforme rede e disponibilidade dos serviços. Uma URL inválida faz parte do exemplo para demonstrar tratamento de falha.

## Comparativo Go, C/pthreads e Python

### Versão Go

```bash
go run ./examples/comparativo/go
```

### Versão C

```bash
cc -std=c11 -Wall -Wextra -Wpedantic -Werror -O2 -pthread \
  examples/comparativo/c/main.c -o /tmp/comparativo-c
/tmp/comparativo-c
```

### Versão Python

```bash
python3 -B examples/comparativo/python/comparativo.py
```

As três versões devem informar `17984` primos entre 2 e 200000.

### Verificação completa

```bash
./examples/comparativo/verificar.sh
```

A verificação:

1. executa os testes Go com o detector de data races;
2. executa os testes Python com `unittest`;
3. executa as versões Go e Python;
4. compila C com avisos tratados como erros;
5. executa a versão C;
6. compara as três saídas byte a byte.

Artefatos de compilação são criados em um diretório temporário e removidos ao final.

## Verificações adicionais do módulo Go

```bash
go test ./...
go test -race ./...
go vet ./...
gofmt -d .
```

O detector de corridas observa apenas os caminhos efetivamente executados. Um resultado sem alertas aumenta a confiança no teste realizado, mas não prova ausência de toda data race possível.

## Verificação isolada da implementação Python

```bash
python3 -B -m unittest discover \
  -s examples/comparativo/python \
  -p 'test_*.py'
```

O exemplo usa somente a biblioteca padrão. Na build padrão do CPython, o GIL limita o paralelismo de bytecode em cargas de CPU; builds *free-threaded* devem ser identificadas e avaliadas separadamente.

## Leitura relacionada

- [Índice da documentação](docs/0_indice.md)
- [Comparativo técnico](docs/3_comparativo_threads.md)
- [Referências bibliográficas](docs/referencias.md)
