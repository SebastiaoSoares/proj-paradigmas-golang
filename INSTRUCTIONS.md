# Instalação e execução

Este guia descreve como reproduzir o CLI e os exemplos do repositório a partir de sua raiz.

## Pré-requisitos

* Go na versão declarada em `go.mod` (1.26.1);
* Python 3.9 ou posterior, disponível como `python3` ou `python`;
* Para o comparativo em C: sistema compatível com POSIX e compilador com suporte a pthreads (`cc`, GCC ou Clang);
* Para a verificação automática do comparativo: shell POSIX e `diff`.

O módulo usa apenas a biblioteca padrão de Go, sem dependências Go de terceiros.

## Compilação e Execução da CLI Unificada

O repositório conta com um menu interativo (CLI). Você pode rodá-lo de duas maneiras:

**Opção A: Execução Direta**
Ideal para testes rápidos e desenvolvimento.
```bash
go run .
```

**Opção B: Compilação do Binário**
Gera o arquivo executável otimizado. Recomendado para uso final.
```bash
# Para compilar:
go build -o app_paradigmas main.go

# Para executar (Linux/macOS):
./app_paradigmas

# Para executar (Windows):
app_paradigmas.exe
```

## Como Navegar no Menu CLI

O menu aceita frases com espaços, limpa a tela entre interações e usa cores ANSI. Para interagir com a aplicação:

1. Observe as funcionalidades listadas na tela (1 - exemplos básicos, 2 - interoperabilidade, 3 - estudo de caso, 4 - comparativo, 0 - encerramento).
2. Digite no terminal o número correspondente à opção desejada.
3. Pressione a tecla **Enter** para confirmar e executar a funcionalidade.
4. Para retornar ou finalizar a aplicação, utilize a opção de saída (geralmente `0`).

A opção **4** primeiro apresenta a execução visual dos workers nas três linguagens e depois executa integralmente a verificação: testes Go com detector de corridas, testes Python, compilação C, exibição das três saídas completas e comparação dos resultados. Ela requer todos os pré-requisitos do comparativo listados abaixo.

**Nota:** Para produzir logs sem cores, defina a variável de ambiente `NO_COLOR`. O monitoramento requer acesso à internet. Códigos HTTP e mensagens podem variar conforme a rede e a disponibilidade dos serviços. A URL inválida faz parte do exemplo e demonstra o tratamento de falhas.

## Comparativo Go, C/pthreads e Python

### Demonstração visual guiada

```bash
./examples/comparativo/demonstrar.sh
```

Para executar automaticamente, sem pausas entre as linguagens:

```bash
./examples/comparativo/demonstrar.sh --auto
```

Cada implementação aceita `--visual` isoladamente. A pausa visual é didática e não representa desempenho.

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

A verificação executa testes Go com o detector de data races, executa testes Python, compila C com avisos tratados como erros e compara as três saídas. Artefatos de compilação são criados em um diretório temporário e removidos ao final.

## Verificações adicionais

```bash
go test ./...
go test -race ./...
go vet ./...
gofmt -d .
```

O detector de corridas observa somente os caminhos executados e não prova a ausência de toda data race possível.

Para testar apenas a implementação Python:

```bash
python3 -B -m unittest discover \
  -s examples/comparativo/python \
  -p 'test_*.py'
```

Na build padrão do CPython, o GIL limita o paralelismo de bytecode em cargas de CPU; builds *free-threaded* devem ser identificadas e avaliadas separadamente.

## Leitura relacionada

* [Índice da documentação](docs/0_indice.md)
* [Comparativo técnico](docs/3_comparativo_threads.md)
* [Referências bibliográficas](docs/referencias.md)
