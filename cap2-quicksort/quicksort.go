package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Define entrada como todos argumentos passados menos o primeiro
	entrada := os.Args[1:]
	fmt.Printf(">>> Entrada: %s\n", entrada)

	// Utiliza a função nativa make para iniciar o slice com o tipo []int
	// E define o tamanho final do slice como len de entrada
	numeros := make([]int, len(entrada))
	fmt.Printf(">>> Numeros: %d\n", numeros)

	for i, v := range entrada {
		// Convertendo entrada para inteiro e definindo em int_entrada
		int_entrada, err := strconv.Atoi(v)

		if err != nil {
			fmt.Printf("'%s' não é um número válido!\n", v)

			os.Exit(1)
		}
		numeros[i] = int_entrada

		fmt.Printf(">>> int_entrada: %d\n", int_entrada)

	}
	fmt.Println(quicksort(numeros))
}

func quicksort(numeros []int) []int {
	// Checando se o tamanho de numeros é menor igual á 1
	if len(numeros) <= 1 {
		return numeros
	}

	// Fazendo uma copia do slice original para a variavel n
	n := make([]int, len(numeros))
	copy(n, numeros)

	// Definir um suporte com um elemento mais ou menos no meio da lista
	indicePivo := len(n) / 2
	pivo := n[indicePivo]

	// Removendo um elemento do slice usando append:
	// Primeiro contamos do pivo até o final slice
	// Depois do inicio até o pivo
	// As reticências no final são para informar que todos os elementos
	// do segundo slice devem ser adicionados ao slice-base
	n = append(n[indicePivo+1:], n[:indicePivo]...)

	menores, maiores := particionar(n, pivo)

	return append(
		append(quicksort(menores), pivo),
		quicksort(maiores)...)

}

func particionar(
	numeros []int,
	pivo int) (menores []int, maiores []int) {

	for _, n := range numeros {
		if n <= pivo {
			menores = append(menores, n)
		} else {
			maiores = append(maiores, n)
		}
	}

	return menores, maiores
}
