package main

import (
	"fmt"
	"os"
)

func main() {
	entrada := os.Args[1:]
	fmt.Printf(">>> Entrada: %s\n", entrada)
	// Utiliza a função nativa make para iniciar o slice com o tipo []int
	numeros := make([]int, len(entrada))
	fmt.Printf(">>> Numeros: %d\n", numeros)
}
