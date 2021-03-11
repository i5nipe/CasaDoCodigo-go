// Diferente de um array slices tem tamanho variável e pode crescer indefinidamente
package main

import "fmt"

func main() {
	// Cochetes vazios para identificar que é um slice
	var a []int
	primos := []int{2, 3, 5, 7, 11, 13}
	nomes := []string{}

	fmt.Println(a, primos, nomes) // [] [2 3 5 7 11 13] []
}
