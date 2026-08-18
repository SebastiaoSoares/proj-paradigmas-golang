# Fundamentos teóricos: goroutines e channels

Go adota concorrência como uma forma de estruturar o programa: várias tarefas
podem progredir de maneira independente, mesmo quando não executam fisicamente
ao mesmo tempo. Paralelismo é a execução simultânea dessas tarefas em mais de um
núcleo. Uma aplicação Go pode ser concorrente em um único núcleo e também pode
executar em paralelo quando há processadores disponíveis.

O modelo é inspirado na ideia de *Communicating Sequential Processes* (CSP): em
vez de fazer várias tarefas alterarem livremente o mesmo estado, o programa pode
transferir valores entre elas por canais tipados. A máxima prática é “não
comunique compartilhando memória; compartilhe memória comunicando”. Isso não
proíbe memória compartilhada, mas torna a comunicação explícita e verificável.

## 1. Como uma goroutine funciona

Uma **goroutine** é uma função em execução concorrente com as demais funções do
programa. A instrução `go f(x)` avalia `f` e seus argumentos na goroutine atual,
cria uma nova goroutine e inicia a chamada independentemente. O chamador não
espera o retorno e valores retornados por `f` são descartados; resultados devem
ser comunicados por channel, memória sincronizada ou outro mecanismo apropriado.

```go
go processar(pedido) // a chamada passa a ser uma nova tarefa concorrente
```

Ela não equivale a uma thread do sistema operacional. Goroutines são gerenciadas
pelo runtime de Go e começam com uma pilha pequena, que cresce e diminui conforme
a necessidade. Por isso, criar milhares delas costuma consumir menos memória e
tempo de criação do que criar o mesmo número de threads.

### 1.1 Escalonador G–M–P

O runtime representa a execução por três entidades:

- **G (goroutine):** guarda o estado da tarefa, como pilha e ponto de execução;
- **M (machine):** corresponde a uma thread do sistema operacional que executa
  código Go;
- **P (processor):** contém os recursos necessários para executar código Go e
  uma fila local de goroutines prontas.

Um M precisa estar associado a um P para executar um G. A quantidade de Ps é
determinada por `GOMAXPROCS` e limita quantas goroutines podem executar código Go
simultaneamente. Esse multiplexamento é frequentemente resumido como **M:N**:
muitas goroutines são distribuídas sobre um conjunto menor de threads.

Em linhas gerais, o ciclo é:

1. a instrução `go` cria um G executável e o coloca em uma fila;
2. um P seleciona um G de sua fila local (ou da fila global);
3. um M associado ao P executa esse G;
4. se a fila local esvaziar, o P pode obter trabalho de outras filas (*work
   stealing*), distribuindo a carga;
5. ao terminar, o G é encerrado; ao bloquear, deixa a execução disponível para
   outra goroutine.

Se uma goroutine bloqueia em I/O de rede integrado ao runtime, o *network
poller* pode estacioná-la enquanto outro G utiliza o M. Se um M fica preso em
uma chamada de sistema, o runtime pode separar seu P e associá-lo a outro M.
Pontos seguros e preempção impedem, em condições normais, que uma tarefa longa
monopolize indefinidamente um P.

### 1.2 Estados, bloqueio e término

Uma goroutine alterna conceitualmente entre estados como executável, em execução
e em espera. Ela pode ser estacionada ao aguardar um channel, mutex, temporizador
ou I/O e voltar à fila quando o evento ocorrer. Isso evita manter uma thread
ocupada apenas esperando.

Não há operação pública para “matar” uma goroutine. A função deve retornar por
conta própria, normalmente após receber cancelamento (por exemplo, via
`context.Context`) ou após o channel de entrada ser fechado. Se `main` retorna,
o programa termina sem esperar automaticamente as outras goroutines. Quando é
necessário aguardar conclusão, usam-se recursos como `sync.WaitGroup`.

## 2. Como um channel funciona

Um **channel** é um conduit tipado: `chan T` transporta somente valores do tipo
`T`. Ele é criado com `make`, pode ser bidirecional ou restringido a apenas envio
(`chan<- T`) ou recebimento (`<-chan T`).

```go
pedidos := make(chan int)    // sem buffer
fila := make(chan int, 10)   // buffer com dez posições

pedidos <- 42                // envio
pedido := <-pedidos          // recebimento
```

Internamente, o runtime mantém para cada channel uma estrutura com tipo dos
elementos, estado de fechamento, filas de goroutines remetentes e receptoras e,
quando solicitado, um buffer circular. Um bloqueio interno protege as alterações
dessa estrutura. Isso é detalhe de implementação; a garantia usada pelo programa
vem da especificação e do modelo de memória, não do formato interno.

### 2.1 Sem buffer e com buffer

Em um channel **sem buffer**, o envio só completa quando há um recebimento
correspondente. A transferência funciona como um encontro (*rendezvous*) e
sincroniza produtor e consumidor.

Em um channel **com buffer**, um envio pode completar enquanto houver espaço e
um recebimento pode completar enquanto houver valores. Se o buffer enche, o
remetente espera; se esvazia, o receptor espera. O buffer desacopla temporariamente
as velocidades, mas não aumenta obrigatoriamente o trabalho executado por segundo.

As operações seguem ordem FIFO para os valores enviados por um mesmo channel. O
escalonador, porém, não promete uma ordem geral de execução entre goroutines.

## 3. Por que a comunicação é segura

A segurança pode ser demonstrada pelas relações de sincronização definidas no
[modelo de memória de Go](https://go.dev/ref/mem):

1. **o envio em um channel acontece antes da conclusão do recebimento
   correspondente**;
2. **o fechamento de um channel acontece antes de um recebimento que retorna o
   valor zero por causa desse fechamento**;
3. para um channel de capacidade `C`, o recebimento número `k` acontece antes da
   conclusão do envio número `k+C`. No caso sem buffer (`C = 0`), o recebimento
   acontece antes da conclusão do envio correspondente.

“Acontece antes” (*happens-before*) significa que os efeitos de memória anteriores
à primeira operação ficam ordenados e visíveis depois da segunda. Considere:

```go
var resultado string
pronto := make(chan struct{})

go func() {
    resultado = "concluído" // A
    close(pronto)           // B
}()

<-pronto                    // C
fmt.Println(resultado)      // D
```

Na goroutine produtora, A ocorre antes de B pela ordem do programa. Pela regra de
fechamento, B acontece antes de C; no consumidor, C ocorre antes de D. Pela
transitividade, A acontece antes de D. Assim, a leitura observa a escrita e não
há *data race* nesse acesso. Não é uma coincidência de temporização: é uma
garantia formal do modelo de memória.

O channel também transfere cada valor como uma unidade: remetentes e receptores
não precisam implementar manualmente uma fila compartilhada, um mutex e uma
condição. Se a mensagem contém ponteiros, mapas ou slices, contudo, os dados
apontados ainda podem ser compartilhados. Alterá-los concorrentemente depois do
envio exige nova sincronização. Channels tornam a **comunicação pelo channel**
segura; eles não tornam automaticamente segura toda memória do programa.

## 4. Fechamento, iteração e `select`

Somente o produtor responsável deve fechar um channel para anunciar que não
haverá novos valores. Enviar ou fechar novamente um channel fechado causa
`panic`. Receber de um channel fechado ainda drena valores armazenados; depois
retorna imediatamente o valor zero e `ok == false`.

```go
valor, ok := <-pedidos
if !ok {
    // não haverá mais pedidos
}

for valor := range pedidos {
    processar(valor) // termina quando o channel é fechado e drenado
}
```

Um channel `nil` bloqueia envios e recebimentos indefinidamente, comportamento
útil para desabilitar dinamicamente casos de um `select`, mas perigoso se não for
intencional.

`select` aguarda várias operações de channel. Se mais de uma está pronta, escolhe
uma de modo pseudoaleatório uniforme; um caso `default` permite continuar sem
bloquear. Com channel de cancelamento ou `context.Context`, ele também impede que
workers fiquem presos e vazem recursos.

```go
select {
case pedido := <-pedidos:
    processar(pedido)
case <-ctx.Done():
    return
}
```

## 5. Segurança e eficiência em conjunto

Channels combinam propriedades que seriam implementadas separadamente com
threads tradicionais:

- **segurança de tipos:** somente valores compatíveis podem ser enviados;
- **exclusão interna:** a fila e as esperas do channel não sofrem corrupção por
  acessos simultâneos;
- **ordenação de memória:** envio, recebimento e fechamento criam relações
  *happens-before*;
- **espera cooperativa:** uma goroutine bloqueada é estacionada pelo runtime, em
  vez de consumir CPU em espera ativa;
- **controle de pressão:** channels sem buffer sincronizam ritmos e buffers
  limitados impõem *backpressure* quando a capacidade acaba;
- **composição:** `select` reúne dados, temporização e cancelamento sem uma fila
  de polling escrita pela aplicação.

Essas vantagens não dispensam projeto cuidadoso. Buffers enormes podem apenas
ocultar um consumidor lento; um envio sem receptor causa bloqueio; ciclos de
espera podem causar *deadlock*; e estado compartilhado fora do channel ainda pode
ter *data race*. Deve-se definir claramente quem produz, quem consome, quem fecha
e como cada goroutine termina. O detector `go test -race ./...` ajuda a encontrar
acessos concorrentes sem sincronização, mas não substitui essas regras de projeto.

## 6. Relação com os exemplos do projeto

Em [`internal/exemplos/channels.go`](../internal/exemplos/channels.go), o channel
sem buffer `pedidos` implementa produtor–consumidor. Cada `pedidos <- pedido`
aguarda o `range` receber o valor; ao final, o produtor fecha o channel e o
consumidor sai do laço depois de drenar todos os pedidos.

Em [`internal/exemplos/waitgroup.go`](../internal/exemplos/waitgroup.go), três
goroutines executam os trabalhos, e o `WaitGroup` impede que a função retorne
antes de todas chamarem `Done`. O exemplo mostra concorrência e sincronização de
término; o exemplo de channels acrescenta comunicação tipada entre tarefas.

## Referências essenciais

Uma bibliografia consolidada do trabalho está em [referencias.md](referencias.md).

---

[Anterior: Contexto histórico e conceitos](1_historia_e_conceitos.md) | [Início](0_indice.md) | [Próximo: Comparativo técnico](3_comparativo_threads.md)
