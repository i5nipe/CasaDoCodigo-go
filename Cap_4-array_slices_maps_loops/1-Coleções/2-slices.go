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

	// Após modificar o slice original o novo também foi alterado
	fmt.Println("Original pós modificação:", original) // [1 2 66 4]
	fmt.Println("Novo pós modificação:", novo)         // [2 66]

	// -------------------{ Função append() }---------------------------
	// função append:
	// func append(slice []Tipo, elementos ...Tipo) []Tipo

	s := make([]int, 0)
	s = append(s, 23)
	fmt.Println(s) // [23]

	// Adiciando um valor no inicio de um slice com uma nova variavel
	s1 := []int{23, 24, 25}
	n1 := []int{22}
	s1 = append(n1, s1...)
	fmt.Println(s1) // [22 23 24 25]

	// Adicionando um valor no inico de um slice com a forma literal do append
	s2 := []int{23, 24, 25}
	s2 = append([]int{22}, s2...)
	fmt.Println(s2) // [22 23 24 25]

	// Adicionando um valor no meio do slice
	s3 := []int{1, 2, 5}
	v := []int{3, 4}
	s3 = append(s3[:2], append(v, s3[2:]...)...)
	fmt.Println(s) // [1 2 3 4 5]

	// Removendo valores de um slice
	s4 := []int{45, 1, 2, 3, 4, 5}
	s4 = s4[1:]
	fmt.Println(s4) // [1 2 3 4 5]

	// Removendo valores do meio de um slice
	s5 := []int{1, 2, 3, 4, 65, 25, 5, 6}
	s5 = append(s5[:4], s5[6:]...)
	fmt.Println(s5) // [1 2 3 4 5 6]

	// ----------------------{ Função copy() }-------------------------
	// func copy(destino, origem []Tipo) int

	// Copiando um slice para outro
	numeros := []int{1, 2, 3, 4, 5}
	dobro := make([]int, len(numeros))

	copy(dobro, numeros)

	for i := range dobro {
		dobro[i] *= 2
	}

	fmt.Println(numeros) // [1 2 3 4 5]
	fmt.Println(dobro)   // [2 4 6 8 10]
}
