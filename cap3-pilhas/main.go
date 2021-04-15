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
	_, err := pilha.Desempilhar()
	if err != nil {
		fmt.Println(err)
	}

}
