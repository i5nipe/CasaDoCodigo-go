package main

import "fmt"

func main() {
	a, b := 0, 10

	// Como a única estrutura de repetição em go é for, podemos colocar uma condição
	// lógica e será executado enquanto for verdadeira, similar ao while de
	// outras linguagens.
	for a < b {
		a += 1
	}
	fmt.Println(a)

	// -----------------

	// Podemos utilizar a syntax do C para fazer uma estrutura de repetição
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// -----------------

	// E também é possivel usar a váriavel fora do looping se declarada antes
	var i2 int
	for i2 = 0; i2 < 10; i2++ {
		// ...
	}
	fmt.Println(i)

	// -----------------

	// Podendo omitir qualquer parte da syntax e deixado o ";"
	// Porém se for utilizada apenas uma condição como no primeiro exemplo o ";" pode ser omitido
	var i3 int
	for i3 = 0; i3 < 10; {
		k += 1
	}

	// -----------------
	numeros := []int{1, 2, 3, 4, 5, 6, 7}
}
