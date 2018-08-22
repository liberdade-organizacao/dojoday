Dojo Day 2018
=============

Descrição do problema
---------------------

```
Inspirado em fatos reais
```

Você foi contactado por uma psicóloga para fazer um programa de computador que confere se a execução de um teste de matrizes progressivas de Raven é válida. O sobrinho dela, estudante de engenharia da computação do Colégio para Estudantes Universitários de Brasília, participante do Dojo de TDD de sua universidade, já começou a fazer o serviço, dividindo a aplicação em camadas e fazendo testes unitários, como recomendado no dojo; mas ele não pode continuar porque conseguiu outro emprego. Sua tarefa então é continuar o trabalho dele, fazendo a classe final, que implementa a lógica do sistema, mantendo as boas práticas de engenharia de software que ele começou.

A saber, um teste de Raven é composto por várias séries. Cada série possui um conjunto de respostas. O resultado de uma série é o número de questões que um sujeito acertou. Uma série é dita válida se o seu resultado não difere, em magnitude, de mais de 2 pontos do esperado. O resultado esperado é dado por uma tabela, que mostra, para o resultado final do teste, qual o resultado esperado de cada série. Um teste é dito válido se todas as suas séries são válidas.

Nesta perspectiva, o sobrinho dividiu a aplicação em três camadas: uma chamada View, que recebe as respostas de um sujeito; outra chamada Controller, que irá checar a validade desta execução se baseando nos valores de validade que podem ser resgatados pela camada Model, a terceira e última camada. Ao analisar o código já escrito, você percebe que as camadas View e Model já foram implementadas, faltando somente a implementação da camada Controller.

Entregas
--------

Você deverá implementar as seguintes funções em Go para completar o programa:

``` go
// Indica se a idade fornecida está de acordo com as restrições do teste
func ValidateAge(age int) bool { }

// Indica se a resposta dada está de acordo com o item atual
func ValidateAnswer(item, answer int) bool { }

// Devolve o número de itens que o teste contém
func GetNumberOfItems() int { }

// Checa se as respostas dadas para um teste estão de acordo com a idade
// do sujeito
func ValidateApplication(age int, answers []int) int { }

// Devolve a descrição do n-ésimo item do teste
func GetNthItem(n int) string { }
```
