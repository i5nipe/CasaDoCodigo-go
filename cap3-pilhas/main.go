package main

import (
	"fmt"
)

func main() {
	pilha := Pilha{}

	// Repare que todos os metodos chamados se iniciam com letra maiúscula
	// até mesmo o println do fmt. O que acontece é que todo identificador
	// que se inicia com letra maiúscula é automaticamente exportado e visível
	// fora do pacote.
	fmt.Println("Pilha criada com tamanho de ", pilha.Tamanho()) // Retorna 0
	fmt.Println("Vazia?", pilha.Vazia())                         // Retorna true

	// Empilhando alguns objetos
	pilha.Empilhar("Go")
	pilha.Empilhar("2009")
	pilha.Empilhar("3.14")
	pilha.Empilhar("Fim")

	fmt.Println("Tamanho após empilhar 4 valores: ", pilha.Tamanho()) // 4
	fmt.Println("Vazia?", pilha.Vazia())                              // false

	// For similar ao while de outras linguagem
	for !pilha.Vazia() {
		// Estamos ignorando a saida de erro usando "_"
		v, _ := pilha.Desempilhar()

		fmt.Println("Desempilando ", v)
		fmt.Println("Tamanho: ", pilha.Tamanho())
		fmt.Println("Vazia?", pilha.Vazia())

	}
	// Agora estamos ignorando o primeiro valor
	_, err := pilha.Desempilhar()
	if err != nil {
		fmt.Println(err)
	}

}

//Definindo o tipo pilha e seus metodos.
// Como Go não é orientado a objetos não possui o conceito de classes
// Ao invés disso definimos estruturas de dados em forma de structs
type Pilha struct {
	// O tipo interface{} é conhecido como interface vazia
	// Se usada como entrada de uma função ela vai aceitar qualquer tipo de objeto.
	// Otimo blog para entender melhor:
	// https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go
	valores []interface{}
}

// Por mais que seja muito semelhante não estamos definindo uma função e sim um método
// A diferença marcante é que metódos definem um objeto receptor
// que deve ser especificado entre parênteses antes do nome do método
// que nesse caso é uma pilha do tipo Pilha.
func (pilha Pilha) Tamanho() int {
	return len(pilha.valores)
}
