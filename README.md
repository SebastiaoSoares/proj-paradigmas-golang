# **Trabalho Prático: Paradigmas de Programação**

**Universidade Federal do Cariri (UFCA)**  \- **Centro de Ciência e Tecnologia (CCT)**

**Curso:** Bacharelado em Engenharia de Software

**Disciplina:** Paradigmas de Programação

**Professor:** Rafael Will Macedo de Araújo

## **Go \- Concorrência Baseada em Goroutines e Channels**

Este repositório contém o Trabalho Prático desenvolvido pela **Equipe 2** para a disciplina de Paradigmas de Programação. O objetivo central deste projeto é investigar a linguagem **Go (Golang)**, com foco especial em seu modelo de concorrência que utiliza *Goroutines* e *Channels*.

Buscamos demonstrar como Go apoia o desenvolvimento de aplicações concorrentes e distribuídas, comparando sua abordagem com **threads POSIX em C** e **threads em Python**. C aproxima o trabalho dos conceitos de sincronização estudados em Sistemas Operacionais; Python acrescenta o contraste com uma API de alto nível e com o GIL do CPython.

### **Equipe 2 (Integrantes)**

1. Elder Rayan Oliveira Silva ([@eldrayan](https://github.com/eldrayan))
2. Espedio Ramom Mascena Ricarto ([@RamomRicarto](https://github.com/RamomRicarto))
3. Manoel Junio Duarte da Silva ([@Junio404](https://github.com/Junio404))
4. Pedro Yan Alcantara Palácio ([@pedropalacioo](https://github.com/pedropalacioo))
5. Sabrina Alencar Soares ([@sabrinaalencaar](https://github.com/sabrinaalencaar))
6. Samuel Wagner Tiburi Silveira ([@samsilveira](https://github.com/samsilveira))
7. Sebastião Sousa Soares ([@SebastiaoSoares](https://github.com/SebastiaoSoares))

## **Objetivos do Projeto**

Atendendo às especificações da disciplina, este projeto cobre os seguintes tópicos:

* **Contextualização Histórica:** Problemas que motivaram a criação do Go no Google.  
* **Fundamentos Teóricos:** Mecanismos de funcionamento de *Goroutines* e *Channels*.  
* **Comparativo Técnico:** Vantagens, limitações e cenários do modelo de concorrência do Go frente às [threads em C e Python](./docs/3_comparativo_threads.md).
* **Análise Crítica:** Avaliação da maturidade, ecossistema, comunidade e limitações da linguagem.  
* **Prática e Estudo de Caso:** Códigos de demonstração, um comparativo visual ao vivo e uma aplicação real (ex: Servidor Web / Processamento Paralelo) desenvolvida em Go.

## **Estrutura do Repositório**

O repositório foi organizado da seguinte forma para facilitar a avaliação e o estudo futuro:
```
├── docs/                                   # Documentação teórica interconectada
│   ├── 0_indice.md                         # Menu central da documentação
│   ├── 1_historia_e_conceitos.md       
│   ├── 2_teoria_goroutines_channels.md 
│   ├── 3_comparativo_threads.md        
│   ├── 4_analise_critica.md            
│   └── referencias.md                  
├── internal/                               # Pacotes usados pelo CLI
│   ├── exemplos/                           # WaitGroup e channels
│   ├── interop/                            # Integração local entre Go e Python
│   ├── estudocaso/                         # Monitoramento concorrente de URLs
│   └── terminal/                           # Identidade visual do terminal
├── examples/comparativo/                   # Mesmo worker pool em três linguagens
│   ├── go/                                 # Goroutines e channels
│   ├── c/                                  # Mutex e variáveis de condição POSIX
│   ├── python/                             # Thread e Queue da biblioteca padrão
│   ├── demonstrar.sh                       # Visualiza fila e workers nas três linguagens
│   └── verificar.sh                        # Testa e compara as três implementações
├── go.mod                                  # Arquivo de módulo gerado pelo Go
├── main.go                                 # Ponto de entrada do CLI unificado
├── instrucoes.md                           # Instalação, execução e verificação
└── README.md                               # Visão geral do projeto
```

## **Como Executar os Códigos**

Para garantir a reprodutibilidade dos nossos experimentos, preparamos um guia passo a passo.

Por favor, consulte o arquivo [**instrucoes.md**](./instrucoes.md) na raiz deste repositório. Lá você encontrará:

* Requisitos de sistema (versão do Go, dependências, etc).  
* Instruções de compilação e execução dos programas.
* Passo a passo para rodar o estudo de caso e o comparativo Go/C/Python.
