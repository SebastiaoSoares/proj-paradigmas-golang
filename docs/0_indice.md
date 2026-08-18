# Índice da documentação

Esta documentação apresenta a linguagem Go com ênfase em concorrência baseada em goroutines e channels.

1. [Contexto histórico e conceitos](1_historia_e_conceitos.md)
2. [Teoria: goroutines e channels](2_teoria_goroutines_channels.md)
3. [Comparativo técnico: Go, threads POSIX em C e threads em Python](3_comparativo_threads.md)
4. [Análise crítica](4_analise_critica.md)
5. [Referências bibliográficas](referencias.md)

## Exemplos práticos

- [Estudo de caso: monitoramento concorrente de URLs](../internal/estudocaso/monitor.go)
- [Worker pool em Go](../examples/comparativo/go/main.go)
- [Worker pool com POSIX threads em C](../examples/comparativo/c/main.c)
- [Worker pool com threads e filas em Python](../examples/comparativo/python/comparativo.py)
- [Verificação automática do comparativo](../examples/comparativo/verificar.sh)

## Como reproduzir

Consulte as [instruções de instalação e execução](../instrucoes.md).
