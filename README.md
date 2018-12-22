# sequence-diagram-gen
## Definições
* Um (arquivo) contém código em uma (linguagem) qualquer.
* Um (módulo) é o conteúdo principal de um (arquivo) que contém: (métodos) e (módulos).
* Um (método) pode possuir (asserções) que definem (retorno) ou (chamada).
* Um (módulo) pode (chamar) outros (módulos).
* Um (módulo) obrigatóriamente deve possuir (métodos).
* Um (módulo) pode possuir (pontos de entrada).
* Um (projeto) pode possuir vários (pontos de entrada).
* Quando um (projeto) for uma (API), os (pontos de entrada) serão as (actions) de uma (controller).
* Uma (action) é um (método) de uma (controller).
* Uma (controller) é um (módulo) localizado no pacote web e/ou com o nome do (arquivo) possuindo o sufixo controller.
* Quando um (projeto) for um (programa), o (ponto de entrada) será a (main).
* A (main) é um (método).
* Uma (árvore de utilização) irá possuir um (ponto de entrada) e irá
armazenar todas as (chamadas) que o (ponto de entrada) faz para outros (módulos) e as (chamadas) que esses (módulos) fazem para outros.
* Uma (visualização) será um meio gráfico de mostrar uma (árvore de utilização)

## Core-domain
* O sequence-diagram-gen irá carregar todos (arquivos) de uma determinada (linguagem) de um (projeto).
* Após carregar todos os (arquivos) os (módulos) serão filtrados.
* Encontrando todos os (módulos), será filtrado aqueles (módulos) que possuem (pontos de entrada).
* A partir dos (pontos de entrada), será buscado todas as (chamadas) para outros (módulos) e (retornos), de forma recursiva. O resultado dessa busca será uma (árvore de utilização).
* A partir das (árvores de utilização), serão construídos (visualizações)
