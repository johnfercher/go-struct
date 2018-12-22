# sequence-diagram-gen
* Um módulo é o conteúdo de um arquivo de texto de código com namespaces, classes, propriedades e métodos
* Um módulo pode chamar outros módulos, se isso ocorrer, obrigatóriamente o módulo chamado deverá ser uma propriedade do módulo pai
* Um classe com comportamento é um módulo
* Um ponto de entrada é um módulo
* Um projeto pode possuir vários pontos de entrada
* Em uma API um ponto de entrada será uma Action de uma Controller
* Em um programa um ponto de entrada será a Main
* Uma (árvore de utilização) irá possuir um (ponto de entrada) e irá
armazenar todas as chamadas que o (ponto de entrada) faz para outros (módulos) e as chamdas que esses (módulos) fazem para outros.
* Uma visualização será um meio gráfico de mostrar uma (árvore de utilização)

* O sequence-diagram-gen irá analisar todos os arquivos de um projeto para encontrar (módulos) para serem análisados.
* Encontrando todos os (módulos), será buscado todos os (pontos de entrada).
* A partir dos (pontos de entrada), será iniciado uma busca que irá construir as (árvores de utilização)
* A partir das (árvores de utilização), será construído uma (visualização) da mesma.