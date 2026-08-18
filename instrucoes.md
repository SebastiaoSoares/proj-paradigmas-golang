# Instalação e execução

Este guia descreve como reproduzir o CLI e os exemplos do repositório a partir de sua raiz.

## Pré-requisitos

* Go 1.26.1 ou posterior, conforme a versão mínima declarada em `go.mod`;
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
go build -o app_paradigmas .

# Para executar (Linux/macOS):
./app_paradigmas

# Para executar (Windows):
app_paradigmas.exe
```

## Como navegar no menu CLI

Após executar `go run .` ou iniciar o binário compilado, o menu apresenta as opções `1`, `2`, `3` e `0`. Digite somente o número desejado e pressione **Enter**.

### Opção 1 — Exemplos básicos

1. Digite `1` e pressione **Enter**.
2. Observe o exemplo de `WaitGroup`, que aguarda a conclusão das goroutines.
3. Em seguida, observe o exemplo de channels no padrão produtor–consumidor.
4. Quando os dois exemplos terminarem, pressione **Enter** para voltar ao menu principal.

### Opção 2 — Interoperabilidade Python

1. Digite `2` e pressione **Enter**.
2. No campo solicitado, digite o texto que deseja converter e pressione **Enter**. O texto pode conter espaços.
3. O programa inicia o Python e exibe o texto convertido em letras maiúsculas.
4. Pressione **Enter** para voltar ao menu principal.

Essa opção requer que `python3` ou `python` esteja disponível no `PATH`.

### Opção 3 — Estudo de caso: monitor de uptime

1. Digite `3` e pressione **Enter**.
2. Aguarde as seis verificações HTTP concorrentes terminarem.
3. Confira o código HTTP das URLs disponíveis e as mensagens das URLs que falharam.
4. Pressione **Enter** para voltar ao menu principal.

Essa opção requer acesso à internet. Códigos HTTP e mensagens podem variar conforme a rede e a disponibilidade dos serviços. A URL `https://invalid-url-for-testing.local` é inválida propositalmente e demonstra o tratamento de falhas.

### Opção 0 — Encerrar

1. Digite `0` e pressione **Enter**.
2. O programa exibe a mensagem de encerramento e finaliza.

Se outro valor for informado, o programa mostra uma mensagem de opção inválida. Pressione **Enter** para voltar ao menu e tentar novamente.

**Nota:** O menu limpa a tela entre interações e usa cores ANSI. Para produzir logs sem cores, defina a variável de ambiente `NO_COLOR` antes da execução.

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
