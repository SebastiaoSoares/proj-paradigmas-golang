# Apresentação — Go: concorrência com goroutines e channels

## Planejamento geral

- **Duração total planejada:** 15 minutos.
- **Faixa aceita pelo professor:** 12 a 17 minutos, seguida de até 5 minutos para perguntas.
- **Formato:** 11 slides, demonstração ao vivo do CLI e encerramento.
- **Objetivo:** explicar o modelo de concorrência de Go, compará-lo com C/pthreads e Python/threads e demonstrar sua aplicação prática.
- **Equipe:** Elder, Espedito, Manoel, Pedro, Sabrina, Samuel e Sebastião.

> O tempo de conteúdo é de 14 minutos, distribuído igualmente em 2 minutos por integrante. O minuto restante é reservado às trocas de apresentador e à abertura da demonstração.

## Estrutura dos slides e roteiro de fala

### Slide 1 — Abertura e problema de pesquisa

**Responsável:** Sebastião Sousa Soares  
**Tempo:** 30 segundos

**Composição visual:**

- título do trabalho e identificação institucional: UFCA, PROGRAD, CCT e Bacharelado em Engenharia de Software;
- disciplina Paradigmas de Programação, professor Rafael Will Macedo de Araújo e identificação da Equipe 2;
- diagrama simples: `tarefas → concorrência → resultado`;
- pergunta central em destaque: “Como Go simplifica a construção de software concorrente?”.

**Roteiro de fala:**

1. Apresentar o tema e a equipe.
2. Explicar que sistemas modernos precisam coordenar diversas tarefas simultâneas.
3. Introduzir o objetivo: compreender goroutines e channels e compará-los com mecanismos equivalentes em C e Python.

### Slide 2 — Contexto histórico e proposta do Go

**Responsável:** Elder Rayan Oliveira Silva  
**Tempo:** 2 minutos

**Composição visual:**

- linha do tempo: início do projeto no Google em 2007 e lançamento do Go 1.0 em 2012;
- três palavras-chave: **simplicidade**, **compilação rápida** e **concorrência**;
- nomes dos criadores: Robert Griesemer, Rob Pike e Ken Thompson.

**Roteiro de fala:**

1. Contextualizar os desafios de grandes sistemas, tempos de compilação e arquiteturas multicore.
2. Explicar que Go combina tipagem estática, compilação nativa e uma linguagem deliberadamente simples.
3. Relacionar o projeto da linguagem às necessidades de redes, serviços e infraestrutura.

### Slide 3 — Goroutines e escalonamento

**Responsável:** Pedro Yan Alcantara Palácio  
**Tempo:** 1 minuto

**Composição visual:**

```go
go executarTarefa()
```

- diagrama: várias goroutines sendo multiplexadas sobre threads do sistema operacional;
- comparação visual entre uma chamada comum e uma chamada iniciada com `go`;
- destaque para `sync.WaitGroup` como mecanismo de espera por conclusão.

**Roteiro de fala:**

1. Definir goroutine como uma função executada concorrentemente.
2. Explicar que o runtime gerencia a execução sobre threads do sistema operacional.
3. Diferenciar concorrência de paralelismo.
4. Mostrar por que `WaitGroup` é necessário quando o fluxo principal precisa aguardar workers.

### Slide 4 — Channels e interoperabilidade

**Responsável:** Pedro Yan Alcantara Palácio  
**Tempo:** 1 minuto

**Composição visual:**

```go
jobs := make(chan Job)
jobs <- tarefa
resultado := <-jobs
```

- fluxo visual: `produtor → channel → consumidor`;
- setas distintas para envio, recebimento e fechamento;
- fluxo da interoperabilidade: `Go → processo Python → saída em maiúsculas`;
- pequeno alerta sobre deadlocks e goroutine leaks.

**Roteiro de fala:**

1. Definir channels como canais tipados de comunicação e sincronização.
2. Explicar envio, recebimento, bloqueio e fechamento.
3. Relacionar o mecanismo ao modelo CSP.
4. Mostrar que a opção 2 do CLI demonstra interoperabilidade real entre Go e Python.
5. Ressaltar que abstrações de alto nível reduzem código, mas não eliminam erros de concorrência.

### Slide 5 — Comparativo técnico: Go, C e Python

**Responsável:** Samuel Wagner Tiburi Silveira  
**Tempo:** 1 minuto e 15 segundos

**Composição visual:**

| Aspecto | Go | C/pthreads | Python/threads |
| --- | --- | --- | --- |
| Unidade concorrente | goroutine | thread POSIX | `Thread` |
| Distribuição de tarefas | channel | fila circular | `Queue` |
| Sincronização | channel e `WaitGroup` | mutex e condições | `Queue` e `join` |
| Gerenciamento | runtime de Go | aplicação e SO | interpretador e SO |
| CPU em paralelo | sim | sim | limitado pelo GIL no CPython padrão |
| Segurança | channels tipados e race detector | invariantes manuais | filas sincronizadas; GIL não substitui locks |
| Produtividade e manutenção | coordenação concisa | maior volume de código | API simples, com limites em CPU-bound |

**Roteiro de fala:**

1. Explicar que as três implementações resolvem exatamente o mesmo problema.
2. Comparar o nível de abstração e a quantidade de coordenação explícita.
3. Esclarecer o efeito do GIL em tarefas CPU-bound no CPython padrão.
4. Explicar que a demonstração visual não é benchmark; desempenho exige medição específica.
5. Evitar declarar uma linguagem como universalmente superior: a escolha depende do domínio e dos requisitos.

### Slide 6 — Códigos práticos e organização em pacotes

**Responsável:** Manoel Junio Duarte da Silva  
**Tempo:** 2 minutos

**Composição visual:**

- árvore resumida dos pacotes `internal/exemplos`, `internal/interop` e `internal/comparativo`;
- trecho curto do exemplo com `WaitGroup`;
- fluxo produtor–consumidor implementado com channel.

**Roteiro de fala:**

1. Explicar como os exemplos práticos foram separados em pacotes reutilizáveis.
2. Mostrar a função dos exemplos de `WaitGroup` e channels.
3. Relacionar a organização dos pacotes à integração posterior pelo menu principal.

### Slide 7 — Estudo de caso: monitor de uptime

**Responsável:** Espedito Ramom Mascena Ricarto
**Tempo:** 2 minutos

**Composição visual:**

```text
lista de URLs → goroutines HTTP → channel de resultados → terminal
```

- captura da opção 3 do CLI;
- destaque para timeout, `WaitGroup`, fechamento do channel e tratamento de falhas;
- uma URL inválida identificada como caso de erro proposital.

**Roteiro de fala:**

1. Apresentar o monitor de uptime como aplicação integrada de concorrência e rede.
2. Explicar que cada URL é consultada concorrentemente com timeout.
3. Mostrar como os resultados chegam por channel e como erros são exibidos sem interromper as demais verificações.

### Slide 8 — Demonstração integrada pelo CLI

**Responsáveis:** Sebastião Sousa Soares (menu e integração) e Samuel Wagner Tiburi Silveira (comparativo)  
**Tempo:** 1 minuto — Sebastião: 15 segundos; Samuel: 45 segundos

**Composição visual:**

- captura ou gravação curta do menu principal;
- comando de execução em destaque:

```bash
NO_COLOR=1 go run .
```

- fluxo da opção 4:

```text
CLI
 ├─ demonstração visual: Go → C → Python
 └─ testes e comparação: verificar.sh → diff → aprovação
```

**Roteiro de demonstração:**

1. Executar o CLI e selecionar a opção `4`.
2. Apontar a distribuição das oito faixas entre quatro workers.
3. Mostrar que a ordem de conclusão pode variar, embora o resultado final seja determinístico.
4. Destacar o total de **17.984 primos** em todas as linguagens.
5. Mostrar as mensagens `Go x C: saídas iguais` e `Go x Python: saídas iguais`.

**Observação:** as pausas do modo visual são didáticas e não representam um benchmark de desempenho.

### Slide 9 — Execução, documentação e reprodutibilidade

**Responsável:** Sabrina Alencar Soares  
**Tempo:** 2 minutos

**Composição visual:**

```bash
go test ./...
go test -race ./...
go vet ./...
./examples/comparativo/verificar.sh
```

- checklist visual com testes Go, testes Python, compilação C e comparação por `diff`;
- selo ou indicador “mesma entrada, mesma saída”;
- links para a documentação em Markdown e para as referências bibliográficas;
- referência ao tratamento de erros e limpeza dos artefatos temporários.

**Roteiro de fala:**

1. Indicar onde estão as instruções unificadas, a navegação e as referências do projeto.
2. Explicar que o script automatiza a validação das três implementações.
3. Destacar que o detector de data races verifica somente os caminhos executados.
4. Identificar as fontes bibliográficas que sustentam as afirmações históricas e técnicas.
5. Relacionar instruções reproduzíveis e saída determinística à facilidade de avaliação do trabalho.

### Slide 10 — Avaliação crítica e aplicações

**Responsável:** Sebastião Sousa Soares  
**Tempo:** 1 minuto

**Composição visual:**

- duas colunas: **vantagens** e **limitações**;
- aplicações sugeridas: APIs, microsserviços, pipelines e infraestrutura;
- riscos: data races, deadlocks, vazamentos de goroutines e complexidade de topologias.

**Roteiro de fala:**

1. Apresentar Go como opção forte para serviços e tarefas concorrentes.
2. Destacar clareza, biblioteca padrão e ferramentas integradas.
3. Reconhecer que channels não substituem projeto cuidadoso de cancelamento, propriedade de dados e encerramento.
4. Reforçar que decisões técnicas devem considerar domínio, equipe e medições reais.

### Slide 11 — Conclusão

**Responsável:** Sebastião Sousa Soares  
**Tempo:** 15 segundos

**Composição visual:**

- uma única mensagem central;
- total validado: **17.984 primos nas três implementações**;
- links ou QR code para o repositório e para a documentação.

**Roteiro de fala:**

> Go não elimina a complexidade da concorrência, mas oferece primitivas concisas e integradas que tornam padrões como worker pools mais claros e produtivos.

Agradecer e abrir espaço para até 5 minutos de perguntas.

## Distribuição consolidada

| Integrante | Responsabilidade principal | Tempo de fala |
| --- | --- | ---: |
| Elder | Contextualização histórica | 2 min |
| Espedito | Estudo de caso integrado | 2 min |
| Manoel | Códigos práticos e pacotes | 2 min |
| Pedro | Fundamentos e interoperabilidade | 2 min |
| Sabrina | Execução, documentação e referências | 2 min |
| Samuel | Comparativo técnico e demonstração | 2 min |
| Sebastião | Abertura, CLI, análise crítica e conclusão | 2 min |
| Transições | Troca de apresentadores e preparação do CLI | 1 min |
| **Total planejado** |  | **15 min** |

## Diretrizes para o material visual

- Usar proporção 16:9 e no máximo duas famílias tipográficas.
- Manter alto contraste e corpo mínimo de 24 pt para textos projetados.
- Preferir diagramas, tabelas curtas e trechos de código de até cinco linhas.
- Não transformar os slides em roteiro escrito; os detalhes ficam neste documento.
- Usar cores consistentes para as linguagens: azul para Go, cinza para C e amarelo para Python.
- Exibir fontes e referências no rodapé quando houver afirmações históricas ou técnicas.
- Evitar animações decorativas; usar movimento apenas para explicar fluxo ou concorrência.

## Plano de contingência

- Manter uma captura de tela ou gravação curta da opção 4 caso falte algum compilador no computador da apresentação.
- Executar previamente `./examples/comparativo/verificar.sh` e guardar a saída esperada.
- Usar `NO_COLOR=1` se o terminal de projeção não renderizar corretamente as cores.
- Se o tempo estiver abaixo de 12 minutos, aprofundar a leitura da tabela comparativa e os limites do race detector.
- Se houver atraso, reduzir detalhes históricos e comentários durante a demonstração, sem remover a conclusão.

## Checklist antes da exposição

- [ ] Slides revisados e exportados também em PDF.
- [ ] CLI e opção 4 testados no equipamento da apresentação.
- [ ] Go, Python, shell POSIX e compilador C disponíveis.
- [ ] Terminal, fonte e zoom ajustados para projeção.
- [ ] Cada integrante ensaiou sua fala com cronômetro.
- [ ] Ensaio completo ficou entre 12 e 17 minutos, preferencialmente próximo de 15 minutos para manter margem.
- [ ] Todos os sete integrantes participam e dominam também as perguntas relacionadas à própria seção.
- [ ] Plano de contingência disponível offline.

---

[Documentação temática](../0_indice.md) | [Comparativo técnico](../3_comparativo_threads.md) | [Instruções de execução](../../INSTRUCTIONS.md)
