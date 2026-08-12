# **Trabalho Prático: Paradigmas de Programação**

**Universidade Federal do Cariri (UFCA)**  \- **Centro de Ciência e Tecnologia (CCT)**

**Curso:** Bacharelado em Engenharia de Software

**Disciplina:** Paradigmas de Programação

**Professor:** \[Nome do Professor\]

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
7. Sebastião Sousa Soares (Gerente de Projeto)

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
├── docs/                                    \# Documentação teórica e pesquisas (Markdown)  
│   ├── 1\_historia\_e\_conceitos.md         \# Contexto histórico e motivações do Go  
│   ├── 2\_teoria\_goroutines\_channels.md   \# Fundamentos de concorrência em Go  
│   ├── 3\_comparativo\_threads.md           \# Comparação técnica (Go vs Threads tradicionais)  
│   ├── 4\_analise\_critica.md               \# Maturidade, ecossistema e limitações  
│   └── referencias.md                       \# Referências bibliográficas  
├── src/                                     \# Código-fonte dos exemplos e estudo de caso  
│   ├── exemplos\_basicos/                   \# Protótipos curtos demonstrando sintaxe  
│   ├── interoperabilidade/                  \# Exemplo de integração (Go \+ Python)  
│   └── estudo\_de\_caso/                    \# Aplicação real concorrente (Estudo principal)  
├── INSTRUCTIONS.md                          \# Guia de compilação, instalação e execução  
└── README.md                                \# Visão geral do projeto (este arquivo)
```

## **Como Executar os Códigos**

Para garantir a reprodutibilidade dos nossos experimentos por qualquer estudante ou avaliador, preparamos um guia passo a passo.

Por favor, consulte o arquivo [**INSTRUCTIONS.md**](./INSTRUCTIONS.md) na raiz deste repositório. Lá você encontrará:

* Requisitos de sistema (versão do Go, dependências, etc).  
* Instruções de compilação e execução para os exemplos básicos.  
* Passo a passo para rodar o Estudo de Caso.

## **Histórico e Organização**

Este projeto foi gerido utilizando o sistema de *Issues* e *Branches* do GitHub, com entregas divididas de forma simultânea entre os 7 membros da equipe, garantindo a conclusão e congelamento do repositório até o prazo estipulado de **18/08/2020**.
