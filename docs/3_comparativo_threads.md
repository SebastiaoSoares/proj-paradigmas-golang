# Comparativo técnico: concorrência em Go, C e Python

## 1. Objetivo e decisão de escopo

Este documento compara diretamente o modelo de concorrência de **Go** com threads em **C/POSIX** e **Python**. C permanece como o contraste principal de baixo nível, enquanto Python acrescenta a perspectiva de uma linguagem de alto nível:

- `pthreads` se relaciona ao conteúdo de threads, escalonamento e sincronização estudado em Sistemas Operacionais;
- C expõe de forma explícita a fila compartilhada, o mutex, as variáveis de condição e o ciclo de vida das threads;
- o contraste evidencia quais responsabilidades são assumidas pelo runtime de Go e quais permanecem com quem desenvolve;
- Python oferece `threading.Thread` e a fila sincronizada `queue.Queue`, mas seu comportamento em tarefas limitadas por CPU exige considerar o GIL da implementação CPython.

Python não substitui C nesta comparação: as três implementações mostram níveis diferentes de abstração sobre concorrência. Também é necessário separar esta implementação adicional do requisito de interoperabilidade. Executar o mesmo algoritmo em três linguagens permite comparação; não constitui, por si só, comunicação entre Go e Python.

## 2. Escopo e método

A análise considera:

- Go conforme sua especificação, seu modelo de memória e seu runtime;
- C com a API **POSIX Threads (pthreads)** definida pelo POSIX.1-2017;
- Python 3.9 ou posterior com `threading` e `queue` da biblioteca padrão, com observações específicas sobre CPython;
- execução em ambiente compatível com POSIX para o exemplo em C;
- um mesmo problema, mesmos dados, quatro workers e a mesma saída nas três implementações.

O exemplo não é um benchmark. Tempo, memória e vazão variam conforme hardware, sistema operacional, compilador, versão do Go, quantidade de workers e carga. Assim, as conclusões quantitativas exigiriam um experimento próprio com repetição, controle de variáveis e análise estatística.

## 3. Conceitos fundamentais

**Concorrência** é a estruturação de tarefas que podem progredir de forma sobreposta. **Paralelismo** é a execução simultânea dessas tarefas, geralmente em mais de um núcleo. Um programa pode ser concorrente sem executar trabalho em paralelo.

Uma **thread** é um fluxo de execução dentro de um processo. Threads do mesmo processo compartilham o espaço de endereçamento, o que facilita a comunicação, mas exige sincronização quando há escrita concorrente.

Uma **goroutine** é uma função executada concorrentemente e gerenciada pelo runtime de Go. Goroutines compartilham o mesmo espaço de endereçamento, mas o runtime as multiplexa sobre threads do sistema operacional. Elas não são “threads do SO menores”: são uma abstração diferente, com pilhas redimensionáveis e escalonamento pelo runtime.

Um **channel** é um canal tipado usado para enviar valores entre goroutines. A comunicação também pode sincronizar a execução. Isso não elimina mutexes nem data races: estado compartilhado ainda deve ser protegido com channels, `sync.Mutex`, operações atômicas ou outra relação de sincronização válida.

Uma **thread Python** criada por `threading.Thread` executa dentro do mesmo processo e compartilha memória com as demais. `queue.Queue` fornece uma fila sincronizada para troca segura de itens. Na configuração padrão do CPython, o **Global Interpreter Lock (GIL)** permite que apenas uma thread execute bytecode Python por vez. Desde o CPython 3.13 também existem builds *free-threaded* capazes de desabilitar o GIL, mas elas não são a configuração padrão. Portanto, “Python não tem paralelismo com threads” não é uma afirmação universal: isso depende da implementação e da build utilizada.

## 4. Comparação direta

| Critério | Go: goroutines e channels | C: POSIX threads | Python: `threading` e `queue` |
| --- | --- | --- | --- |
| Unidade concorrente | Goroutine criada com `go` | Thread criada com `pthread_create` | Thread criada com `Thread.start()` |
| Gerenciamento | Runtime de Go | Programa, biblioteca pthread e sistema operacional | Interpretador, biblioteca padrão e sistema operacional |
| Escalonamento | Runtime multiplexa goroutines sobre threads do SO | A aplicação cria threads; seu escalonamento efetivo é realizado pelo sistema | Threads nativas, condicionadas também pelo GIL quando habilitado no CPython |
| Pilha e criação | Pilha inicial pequena e redimensionável; criação geralmente barata | Recursos e atributos devem ser considerados explicitamente | Custo de thread nativa; API de alto nível reduz o código de gerenciamento |
| Comunicação idiomática | Channels tipados; memória compartilhada também é possível | Memória compartilhada, mutexes e variáveis de condição | `queue.Queue`, locks e memória compartilhada |
| Espera por conclusão | `sync.WaitGroup` ou composição com channels | `pthread_join` para threads juntáveis | `Queue.join()` para tarefas e `Thread.join()` para threads |
| Fila produtor–consumidor | Channel oferece envio, recepção, bloqueio e buffer opcional | A aplicação implementa a fila e a coordena com mutex e condições | `Queue` já fornece buffer e sincronização entre threads |
| Segurança de memória | Tipagem, garbage collector e modelo de memória de Go | Alocação, tempo de vida e acesso ficam sob responsabilidade do programa | Gerenciamento automático de memória, mas estado lógico ainda requer sincronização |
| Data races | Continuam possíveis; detector integrado ao comando `go` | Continuam possíveis; detecção depende de ferramentas adicionais | O GIL não torna operações compostas nem algoritmos automaticamente seguros |
| CPU-bound | Pode executar em paralelo em múltiplos núcleos | Pode executar em paralelo em múltiplos núcleos | GIL limita bytecode na build padrão do CPython; builds *free-threaded* mudam esse cenário |
| Falhas típicas | Data race, deadlock, envio em channel fechado e vazamento de goroutine | Data race, deadlock, erro de mutex/condição, ciclo de vida ou memória | Deadlock, tarefa sem `task_done`, thread não encerrada e contenção no GIL |
| Portabilidade do exemplo | Toolchain Go nas plataformas suportadas | Requer pthreads; não faz parte do padrão ISO C | Biblioteca padrão multiplataforma, sujeita às características do interpretador |
| Controle de baixo nível | Menor e mediado pelo runtime | Maior controle sobre atributos e primitivas | Menor; abstrações do interpretador e da biblioteca padrão |

### 4.1 Criação e ciclo de vida

Em Go, `go f()` inicia `f` em uma nova goroutine e a execução chamadora continua sem aguardar seu término. No exemplo, um `sync.WaitGroup` registra os workers ativos; quando todos terminam, uma goroutine fecha o channel de resultados.

Em C, `pthread_create` recebe a função inicial e um ponteiro de argumento. Cada thread criada como juntável deve ser coletada com `pthread_join`, que também permite ao fluxo principal saber que suas escritas terminaram. Os códigos de retorno da API precisam ser verificados pelo programa.

Em Python, `Thread.start()` inicia cada worker e `Thread.join()` aguarda seu encerramento. Uma sentinela `None` por worker informa que não existem novos jobs. `Queue.join()` tem outra finalidade: aguarda até que cada item inserido tenha uma chamada correspondente a `task_done()`.

### 4.2 Comunicação e sincronização

O exemplo em Go usa dois channels:

1. `jobsCh` distribui faixas numéricas aos workers;
2. `resultsCh` devolve resultados ao agregador.

No exemplo em C, o programa implementa uma fila circular limitada. Um mutex protege `head`, `tail`, `count` e `closed`. A condição `not_empty` suspende consumidores quando a fila está vazia; `not_full` suspende o produtor quando ela está cheia. `pthread_cond_wait` libera o mutex atomicamente durante a espera e volta a adquiri-lo antes de retornar. A condição é testada em um laço `while`, pois acordar não significa que o predicado já é verdadeiro para aquela thread.

No exemplo em Python, `queue.Queue` encapsula a fila limitada e a sincronização necessária. Uma segunda fila transporta os resultados. A implementação é mais próxima visualmente da versão Go, mas cria threads, não goroutines; a semelhança da API não implica o mesmo modelo de escalonamento.

### 4.3 Escalonamento e paralelismo

O runtime de Go pode manter muitas goroutines prontas e distribuí-las entre threads do sistema. A quantidade de threads que executam código Go simultaneamente é influenciada por `GOMAXPROCS`; operações bloqueantes e chamadas ao sistema também afetam o escalonamento.

No programa C, quatro chamadas a `pthread_create` criam quatro threads concorrentes. O programa controla essa quantidade diretamente, enquanto o sistema operacional decide quando cada thread executa. Nenhuma das abordagens garante ganho de desempenho apenas por adicionar workers: sincronização, granularidade das tarefas e quantidade de núcleos podem tornar a versão concorrente mais lenta.

No CPython padrão, quatro threads podem avançar concorrentemente, mas apenas uma executa bytecode Python por vez devido ao GIL. Threads ainda são adequadas para muitas tarefas de entrada/saída, pois o GIL é liberado durante operações bloqueantes. Para esta carga limitada por CPU, não se deve esperar aceleração proporcional na build padrão. Uma build *free-threaded* pode executar threads em paralelo, mas deve ser identificada explicitamente e medida separadamente.

## 5. Exemplo prático reproduzível

As três implementações contam números primos entre 2 e 200000. O intervalo é dividido em oito jobs e processado por quatro workers. A tarefa é limitada por CPU e foi escolhida porque:

- não depende de rede, arquivos externos nem bibliotecas de terceiros;
- produz um resultado determinístico;
- permite observar distribuição de trabalho e sincronização;
- mantém o foco no modelo concorrente, não no domínio da aplicação.

Arquivos:

- [implementação com goroutines e channels](../examples/comparativo/go/main.go);
- [testes da implementação Go](../examples/comparativo/go/main_test.go);
- [implementação com pthreads](../examples/comparativo/c/main.c);
- [implementação com threads e filas em Python](../examples/comparativo/python/comparativo.py);
- [testes da implementação Python](../examples/comparativo/python/test_comparativo.py);
- [verificação automatizada das três versões](../examples/comparativo/verificar.sh).

Fluxo lógico comum:

```text
produtor -> fila de jobs -> 4 workers -> resultados ordenados -> total
```

A diferença está na realização da fila: Go fornece channels como primitiva da linguagem; Python fornece `Queue` em sua biblioteca padrão; em C, a fila circular e seu protocolo de bloqueio fazem parte do código da aplicação.

### Execução

Na raiz do repositório:

```bash
go run ./examples/comparativo/go
```

```bash
python3 -B examples/comparativo/python/comparativo.py
```

```bash
cc -std=c11 -Wall -Wextra -Wpedantic -Werror -O2 -pthread \
  examples/comparativo/c/main.c -o /tmp/comparativo-c
/tmp/comparativo-c
```

Ou execute toda a verificação:

```bash
./examples/comparativo/verificar.sh
```

O script executa os testes Go com o detector de corridas habilitado, os testes Python com `unittest`, compila C tratando avisos como erros e compara as três saídas. O resultado esperado é um total de **17984 números primos**. A ordem em que os workers finalizam pode variar, mas todas as implementações ordenam os resultados pelo identificador da faixa para manter a saída estável.

## 6. Vantagens e limitações

### 6.1 Go

**Vantagens:**

- criação e coordenação concisas para grandes quantidades de tarefas concorrentes;
- channels tipados expressam transferência de dados e sincronização na mesma abstração;
- runtime administra pilhas e multiplexação sobre threads;
- biblioteca padrão oferece `sync`, `context`, temporizadores, rede e ferramentas integradas;
- detector de corridas integrado ao comando `go` facilita testes instrumentados.

**Limitações:**

- abstrações de alto nível não impedem data races ou deadlocks;
- uma goroutine bloqueada sem caminho de cancelamento pode permanecer viva indefinidamente;
- channels em excesso podem ocultar o fluxo de controle e dificultar manutenção;
- o runtime e o garbage collector reduzem o controle direto sobre escalonamento e memória;
- interoperar com bibliotecas C por `cgo` adiciona fronteiras de desempenho, build e segurança.

### 6.2 C com pthreads

**Vantagens:**

- controle explícito sobre threads, atributos e estruturas de sincronização;
- integração natural com software de sistemas e bibliotecas nativas existentes;
- ausência de runtime com garbage collector;
- adequado quando requisitos de plataforma e recursos precisam de controle fino.

**Limitações:**

- mais código para expressar filas e protocolos de encerramento;
- gerenciamento manual de memória e ciclo de vida amplia a superfície de falhas;
- invariantes de mutex e condições não são verificadas pelo sistema de tipos;
- composição de muitos estágios concorrentes tende a exigir coordenação mais complexa;
- portabilidade depende das APIs de threads disponíveis na plataforma.

### 6.3 Python com `threading` e `queue`

**Vantagens:**

- API de alto nível e pouco código para criar threads e filas sincronizadas;
- `Queue` encapsula bloqueio, buffer e coordenação produtor–consumidor;
- adequado para automação e muitas cargas limitadas por entrada/saída;
- garbage collector e gerenciamento automático de memória reduzem riscos presentes em C;
- biblioteca padrão multiplataforma e testes com `unittest` sem dependências externas.

**Limitações:**

- na build padrão do CPython, o GIL limita o paralelismo de bytecode em tarefas de CPU;
- builds *free-threaded* precisam ser avaliadas quanto a disponibilidade, compatibilidade e desempenho;
- o GIL não substitui sincronização para invariantes compartilhadas;
- threads têm custo maior que goroutines e não são indicadas para quantidades extremamente altas;
- desempenho e semântica podem variar entre implementações e builds de Python.

## 7. Cenários ideais de aplicação

| Cenário | Escolha que tende a se adequar melhor | Justificativa |
| --- | --- | --- |
| Servidores de rede, APIs e microsserviços com muitas operações concorrentes | Go | Goroutines, channels, `context` e biblioteca de rede favorecem concorrência estruturada de alto nível |
| Pipelines, worker pools e CLIs concorrentes | Go | Pouco código de coordenação e ferramentas integradas |
| Automação e aplicações com quantidade moderada de I/O bloqueante | Python | `threading` e `Queue` oferecem uma solução simples, especialmente quando o ecossistema Python já é necessário |
| Processamento de CPU em CPython com GIL | Go, C ou processos Python | Threads Python não oferecem paralelismo de bytecode nessa configuração |
| Sistemas embarcados ou componentes próximos do SO com requisitos específicos | C/pthreads | Controle direto e integração com APIs nativas, quando pthreads está disponível |
| Extensão de uma base C que já usa POSIX | C/pthreads | Evita introduzir novo runtime e fronteira de interoperabilidade sem necessidade |
| Latência e uso de memória rigidamente controlados | Depende de medição | C oferece mais controle; Go pode atender muitos casos, mas a decisão exige orçamento e experimento no ambiente alvo |
| Grande quantidade de tarefas independentes e bloqueantes | Go | O runtime foi projetado para multiplexar muitas goroutines |

Essas recomendações não substituem medição. Domínio, experiência da equipe, manutenção, bibliotecas, plataforma e requisitos não funcionais podem pesar mais do que o mecanismo de concorrência isolado.

## 8. Conclusão crítica

Go não “resolve automaticamente” concorrência segura, nem C/pthreads é tecnicamente inferior. A diferença central é de **nível de abstração e distribuição de responsabilidades**.

Go incorpora goroutines e channels à linguagem e delega ao runtime pilhas, multiplexação e parte do ciclo operacional. Isso reduz o código necessário para padrões como worker pools e costuma favorecer clareza em serviços concorrentes. Em contrapartida, o programador ainda precisa definir propriedade dos dados, encerramento, cancelamento e sincronização corretos.

C com pthreads torna mecanismos e estado compartilhado visíveis. Essa explicitude oferece controle e é valiosa em software de sistemas, mas exige que a aplicação implemente e preserve mais invariantes. Python ocupa uma posição intermediária na ergonomia: `Queue` e `Thread` reduzem o código de coordenação, mas o GIL da build padrão do CPython altera o potencial de paralelismo em cargas de CPU.

Portanto, para o contexto deste trabalho e para aplicações com muitas atividades concorrentes, Go apresenta uma abstração produtiva com paralelismo multicore integrado ao runtime. C/pthreads permanece apropriado quando integração nativa, restrições de plataforma ou controle de baixo nível são requisitos predominantes. Python é uma alternativa acessível para I/O concorrente, automação e ecossistemas já baseados na linguagem; para CPU-bound, a configuração do interpretador deve fazer parte explícita da decisão.

## 9. Referências

As fontes primárias utilizadas estão catalogadas em [Referências bibliográficas](referencias.md#comparativo-go-c-e-python).

---

[Anterior: Goroutines e channels](2_teoria_goroutines_channels.md) | [Índice](0_indice.md) | [Próximo: Análise crítica](4_analise_critica.md)
