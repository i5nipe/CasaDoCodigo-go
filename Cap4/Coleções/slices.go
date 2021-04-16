// Diferente de um array slices tem tamanho variável e pode crescer indefinidamente
package main

import "fmt"

func main() {
	// Cochetes vazios para identificar que é um slice
	var a []int
	primos := []int{2, 3, 5, 7, 11, 13}
	nomes := []string{}

	fmt.Println(a, primos, nomes) // [] [2 3 5 7 11 13] []

	// Tamanho e capacidade são duas propriedades importantes em slices e arrays.
	// make() é uma função capaz também de criar um slice.

	// Assinatura da função make: func make([]T, len, cap) [T]
	// T representa o tipo, len o tamanho, cap o tamanho total ou seja a capacidade.

	// Podemos omitir o último argumento, Go assume por padrão o mesmo valor do len
	b := make([]int, 3)
	fmt.Println(b, len(b), cap(b)) // [0 0 0] 3 3

	c := make([]int, 3, 20)
	fmt.Println(c, len(c), cap(c)) // [0 0 0] 3 20
}
