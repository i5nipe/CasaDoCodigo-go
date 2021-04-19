package main

import (
	"fmt"
	"sort"
)

func main() {
	// Declarando quadrados como um map com chave int e
	// valores int com capacidade 15.
	quadrados := make(map[int]int, 15)

	for n := 1; n <= 15; n++ {
		// Atribuindo a quadrados os quadrados dos números de 1 á 15.
		quadrados[n] = n * n
	}

	numeros := make([]int, 0, len(quadrados))

	// Quando usado o range em um map ele vai retornar a chave e o valor
	for n := range quadrados {
		numeros = append(numeros, n)
	}

	// Função Ints() organiza inteiros
	sort.Ints(numeros)

	for _, numero := range numeros {
		quadrado := quadrados[numero]
		fmt.Printf("%d^2 = %d\n", numero, quadrado)
	}
}
