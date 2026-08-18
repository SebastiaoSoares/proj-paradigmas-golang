# Contexto Histórico e Surgimento da Linguagem Go

## O Cenário Tecnológico e a Motivação no Google

Em 2007, quando Robert Griesemer, Rob Pike e Ken Thompson começaram a esboçar o projeto inicial da linguagem Go nos escritórios do Google, o cenário de desenvolvimento de software era bem diferente do que vemos hoje. Sistemas de grande porte em produção eram escritos predominantemente em C++ ou Java. Ao mesmo tempo, os computadores começavam a evoluir para arquiteturas multicore (com vários núcleos de processamento), mas as linguagens dominantes da época ofereciam pouco e complicado suporte para aproveitar isso com segurança e eficiência.

Nesse período, os engenheiros do Google enfrentavam frustrações crescentes para manter e compilar projetos de software em larga escala. Entre os principais problemas que motivaram a criação da nova linguagem, estavam:

* **Complexidade excessiva e burocracia:** projetos gigantescos exigiam uma sobrecarga enorme de gerência de dependências, tipos hierárquicos complexos e sistemas de build demorados, todos herdados das linguagens orientadas a objetos tradicionais da época.
* **Tempos de compilação extremamente longos:** os desenvolvedores perdiam muito tempo esperando a compilação e a linkagem de binários em C++, o que reduzia drasticamente o ritmo de trabalho.
* **O paradoxo da escolha (eficiência vs. produtividade):** não existia uma ferramenta completa. O programador precisava escolher entre compilação eficiente, execução eficiente ou facilidade de programação. Por isso, muitos abriam mão do controle e da segurança das linguagens fortemente tipadas em troca de linguagens dinâmicas e interpretadas, como Python e JavaScript, só para conseguir programar com mais fluidez.
* **Falta de suporte a multiprocessamento e rede:** a computação em nuvem, o processamento em rede e a execução concorrente já eram premissas básicas para o Google, mas programar e gerenciar memória para threads concorrentes nessas linguagens legadas era complexo e propenso a falhas graves.

## Objetivos Originais de Concepção da Linguagem

Diante dessas dificuldades, a solução não foi criar mais bibliotecas e ferramentas para lidar com os problemas antigos, mas sim dar um passo atrás e conceber uma linguagem do zero, pensada para os novos tempos da engenharia de software em larga escala. Os objetivos originais do projeto Go foram:

1. **Unir o melhor dos dois mundos:** combinar a facilidade, a fluidez e a expressividade das linguagens dinamicamente tipadas e interpretadas com a eficiência e a segurança das linguagens compiladas e estaticamente tipadas.
2. **Suporte nativo a concorrência:** trazer o processamento paralelo e concorrente diretamente para as primitivas da linguagem, através das famosas goroutines e channels (fortemente inspiradas no modelo CSP de Tony Hoare), livrando o desenvolvedor de lidar o tempo todo com threads nativas pesadas, mutexes de baixo nível e bloqueios de acesso.
3. **Segurança e gerenciamento automático de memória:** tornar possível escrever código concorrente em grande escala sem o risco constante do acesso manual e descontrolado à memória, implementando para isso um coletor de lixo (garbage collector) leve e eficiente.
4. **Rapidez extrema na compilação:** garantir builds super rápidos, a ponto de compilar um executável grande em um único computador levar no máximo poucos segundos.
5. **Simplicidade e redução de complexidade (abordagem zen):** apostar num design enxuto e conciso, eliminando elementos comuns, porém propensos a bagunçar a linguagem, como declarações adiantadas (forward declarations), herança hierárquica por classes, sintaxe rebuscada, tratamento de exceções baseado em blocos try-catch e arquivos de cabeçalho (headers). As dependências ficariam mais fáceis de gerenciar, priorizando sistemas de tipos composicionais em vez de hierárquicos.
6. **Ecossistema Amigável e *Tooling* Integrado:** Auxiliar fortemente o trabalho do programador reduzindo os debates estilísticos de escrita de código ao fornecer ferramentas embutidas, sendo a principal delas o formatador automático padrão da linguagem, o `gofmt`.

## Referências e Fontes

* [Frequently Asked Questions (FAQ) - The Go Programming Language](https://go.dev/doc/faq)
* [Por que Go não alcançou o estrelato? A história de uma linguagem quase revolucionária - DIO](https://www.dio.me/articles/por-que-go-nao-alcancou-o-estrelato-a-historia-de-uma-linguagem-quase-revolucionaria)
* [Go (linguagem de programação) - Wikipédia](https://pt.wikipedia.org/wiki/Go_(linguagem_de_programação)#História)
