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
	//...
	//...
	// ------------------- Fatiando slices -------------------------
	//...
	fib := []int{1, 2, 3, 4}
	fmt.Println(fib) // [1 2 3 4]

	// Cortou do primeiro até o terceiro
	fmt.Println(fib[:3]) // [1 2 3]

	// Cortou do segundo até o final
	fmt.Println(fib[2:]) // [3 4]

	// Mostra todo o slice
	fmt.Println(fib[:]) // [1 2 3 4]

	// Quando um slice é criado um array é alocado internamente.
	// Quando fatiamos este slice para uma nova variavel o slice vai utilizar
	// do mesmo array interno ou seja quando modificamos a variavel original
	// a nova vai ser alterada também.

	original := []int{1, 2, 3, 4}
	fmt.Println(original) // [1 2 3 4]

	novo := original[1:3]
	fmt.Println(novo) // [2 3]

	original[2] = 66

	fmt.Println("Original pós modificação:", original) // [1 2 66 4]
	fmt.Println("Novo pós modificação:", novo)         // [2 66]

}
