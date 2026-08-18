# Roteiro da apresentação (12–15 minutos)

[Índice](0_indice.md) | [Execução](../INSTRUCTIONS.md) | [Referências](referencias.md)

Este arquivo funciona como material visual navegável. Use os títulos como telas e mantenha o CLI pronto em outro terminal para a demonstração.

## 1. Problema e motivação — Elder (1 min 30 s)

- Sistemas modernos precisam coordenar muitas tarefas de rede e processamento.
- Threads tradicionais oferecem controle, mas aumentam o trabalho de sincronização.
- Pergunta central: como goroutines e channels tornam a concorrência mais simples?

## 2. Origem e proposta do Go — Espedio (1 min 30 s)

- Criado no Google a partir de 2007; Go 1.0 foi lançado em 2012.
- Compilação rápida, tipagem estática, simplicidade e suporte nativo à concorrência.
- Modelo inspirado em *Communicating Sequential Processes* (CSP).

## 3. Goroutines — Manoel (1 min 30 s)

```go
go executarTarefa()
```

- Funções executadas concorrentemente e gerenciadas pelo runtime de Go.
- São multiplexadas sobre threads do sistema operacional.
- `WaitGroup` permite aguardar a conclusão de um conjunto de tarefas.

## 4. Channels — Pedro (1 min 30 s)

```go
pedidos := make(chan int)
pedidos <- 1
pedido := <-pedidos
```

- Channels tipados comunicam valores e também sincronizam goroutines.
- Fechar o channel sinaliza que não haverá novos valores.
- Uso incorreto ainda pode causar bloqueios, deadlocks e vazamentos de goroutines.

## 5. Comparativo Go, C e Python — Sabrina (2 min)

| Aspecto | Go | C/pthreads | Python/threads |
| --- | --- | --- | --- |
| Unidade | goroutine | thread POSIX | `Thread` |
| Comunicação | channel | fila + mutex + condições | `Queue` |
| Gerência | runtime de Go | aplicação/SO | interpretador/SO |
| CPU em paralelo | sim | sim | limitado pelo GIL no CPython padrão |

As três versões contam **17.984 primos** entre 2 e 200.000 usando quatro workers. O resultado é igual; o nível de abstração e a quantidade de coordenação explícita mudam.

## 6. Demonstração do CLI — Samuel (2 min 30 s)

```bash
NO_COLOR=1 go run .
```

1. Opção 1: observar `WaitGroup` e produtor–consumidor com channel.
2. Opção 2: digitar `Olá, equipe!` e mostrar Go chamando Python.
3. Opção 3: mostrar requisições HTTP concorrentes e uma falha tratada.
4. Opção 4: testar integralmente Go, C e Python e conferir que as saídas são iguais.

Se a rede estiver indisponível, explique o timeout e prossiga: a falha é parte do tratamento demonstrado.

## 7. Avaliação crítica e conclusão — Sebastião (2 min)

- Go reduz código de coordenação e se destaca em serviços, pipelines e ferramentas de infraestrutura.
- Concorrência não implica paralelismo e não elimina data races ou deadlocks.
- C favorece controle de baixo nível; Python é simples para I/O, mas o GIL padrão limita CPU-bound.
- Conclusão: Go oferece um equilíbrio forte entre desempenho, clareza e produtividade para aplicações concorrentes.

## Encerramento e perguntas — todos (1 min)

- Elder retoma a pergunta central; Sebastião apresenta a conclusão em uma frase.
- Cada integrante responde primeiro sobre a seção que apresentou.
- Tempo-base: **13 min 30 s**, deixando até 1 min 30 s de margem.

## Checklist antes da apresentação

- Executar `go test -race ./...`, `go vet ./...` e `./examples/comparativo/verificar.sh`.
- Confirmar Go, Python, compilador C e acesso à internet.
- Abrir este roteiro, o CLI e o comparativo antes de iniciar.
- Ensaiar uma vez com cronômetro e manter cada troca de apresentador curta.
