package main

import "fmt"

func main() {
	// Sempre na declaração de um array você deve especificar o seu tamanho.
	// Quando declarada e não inicializada a variável vai conter o zero value
	var a [3]int

	numeros := [5]int{1, 2, 3, 4, 5}

	// As reticências podem ser usadas fazendo assim o compilador calcular o
	// tamanho do array com base na quantidade de elementos declarados
	primos := [...]int{2, 3, 5, 7, 11, 13}

	nomes := [2]string{}

	fmt.Printf("a: %v %T\nnumeros: %v %T\nprimos: %v %T\n nomes: %v %T",
		a, a, numeros, numeros, primos, primos, nomes, nomes)

	// È possivel calcular o tamanho dos Arrays com len()
	fmt.Println(len(a), len(numeros), len(primos), len(nomes))

	//Arrays multidimensionais, arrays com valores de arrays

	var multiA [2][2]int
	multiA[0][0], multiA[0][1] = 3, 5
	multiA[1][0], multiA[1][1] = 7, -2

	multiB := [2][2]int{{2, 13}, {-1, 6}}

	fmt.Println("Multi A:", multiA) // Multi A: [[3 5] [7 -2]]
	fmt.Println("Multi B:", multiB) // Multi B: [[2 13] [-1 6]]
}
