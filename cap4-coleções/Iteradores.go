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

	//
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//
	var i2 int
	for i2 := 0; i2 < 10; i2++ {
		// ...
	}
	fmt.Println(i)
}
