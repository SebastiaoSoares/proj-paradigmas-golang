# Instalação e execução

Este guia descreve como reproduzir o CLI e os exemplos do repositório a partir de sua raiz.

## Pré-requisitos

- Go na versão declarada em [`go.mod`](go.mod);
- Python 3.9 ou posterior, disponível como `python3` ou `python`;
- para o comparativo em C: sistema compatível com POSIX e compilador com suporte a pthreads (`cc`, GCC ou Clang);
- para a verificação automática do comparativo: shell POSIX e `diff`.

O módulo usa apenas a biblioteca padrão de Go, sem dependências Go de terceiros.

## CLI unificada

```bash
go run .
```

O menu oferece:

1. exemplos básicos de WaitGroup e channels;
2. interoperabilidade entre Go e Python;
3. estudo de caso de monitoramento concorrente de uptime;
0. encerramento do programa.

O menu aceita frases com espaços, limpa a tela entre interações e usa cores ANSI.
Para produzir logs sem cores, defina a variável de ambiente `NO_COLOR`.

O monitor requer acesso à internet. Códigos HTTP e mensagens podem variar conforme
a rede e a disponibilidade dos serviços. A URL inválida faz parte do exemplo e
demonstra o tratamento de falhas.

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

A verificação executa testes Go com o detector de data races, executa testes
Python, compila C com avisos tratados como erros e compara as três saídas.
Artefatos de compilação são criados em um diretório temporário e removidos ao final.

## Verificações adicionais

```bash
go test ./...
go test -race ./...
go vet ./...
gofmt -d .
```

O detector de corridas observa somente os caminhos executados e não prova a
ausência de toda data race possível.

Para testar apenas a implementação Python:

```bash
python3 -B -m unittest discover \
  -s examples/comparativo/python \
  -p 'test_*.py'
```

Na build padrão do CPython, o GIL limita o paralelismo de bytecode em cargas de
CPU; builds *free-threaded* devem ser identificadas e avaliadas separadamente.

## Leitura relacionada

- [Índice da documentação](docs/0_indice.md)
- [Comparativo técnico](docs/3_comparativo_threads.md)
- [Referências bibliográficas](docs/referencias.md)
