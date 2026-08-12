# **Trabalho Prático: Paradigmas de Programação**

**Universidade Federal do Cariri (UFCA)**  \- **Centro de Ciência e Tecnologia (CCT)**

**Curso:** Bacharelado em Engenharia de Software

**Disciplina:** Paradigmas de Programação

**Professor:** Rafael Will Macedo de Araújo

## **Go \- Concorrência Baseada em Goroutines e Channels**

Este repositório contém o Trabalho Prático desenvolvido pela **Equipe 2** para a disciplina de Paradigmas de Programação. O objetivo central deste projeto é investigar a linguagem **Go (Golang)**, com foco especial em seu modelo de concorrência que utiliza *Goroutines* e *Channels*.

Buscamos demonstrar como Go simplifica o desenvolvimento de aplicações concorrentes e distribuídas de forma segura e eficiente, comparando sua abordagem com o uso tradicional de *threads* em linguagens como Python, Java e C++.

### **Equipe 2 (Integrantes)**

1. Elder Rayan Oliveira Silva  
2. Espedio Ramom Mascena Ricarto  
3. Manoel Junio Duarte da Silva  
4. Pedro Yan Alcantara Palácio  
5. Sabrina Alencar Soares  
6. Samuel Wagner Tiburi Silveira  
7. Sebastião Sousa Soares

## **Objetivos do Projeto**

Atendendo às especificações da disciplina, este projeto cobre os seguintes tópicos:

* **Contextualização Histórica:** Problemas que motivaram a criação do Go no Google.  
* **Fundamentos Teóricos:** Mecanismos de funcionamento de *Goroutines* e *Channels*.  
* **Comparativo Técnico:** Vantagens e limitações do modelo de concorrência do Go frente às *threads* tradicionais.  
* **Análise Crítica:** Avaliação da maturidade, ecossistema, comunidade e limitações da linguagem.  
* **Prática e Estudo de Caso:** Códigos de demonstração e uma aplicação real (ex: Servidor Web / Processamento Paralelo) desenvolvida em Go.

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
├── internal/                               # Código-fonte organizado em pacotes Go
│   ├── exemplos/
│   ├── interoperabilidade/
│   └── estudocaso/
├── go.mod                                  # Arquivo de módulo gerado pelo Go
├── main.go                                 # Ponto de entrada com o menu interativo
├── instrucoes.md                           # Como rodar o CLI e usar o menu
└── README.md                               # Visão geral do projeto
```

## **Como Executar os Códigos**

Para garantir a reprodutibilidade dos nossos experimentos, preparamos um guia passo a passo.

Por favor, consulte o arquivo [**INSTRUCTIONS.md**](./INSTRUCTIONS.md) na raiz deste repositório. Lá você encontrará:

* Requisitos de sistema (versão do Go, dependências, etc).  
* Instruções de compilação e execução do menu CLI.  
* Passo a passo para rodar o Estudo de Caso e exemplos básicos.
