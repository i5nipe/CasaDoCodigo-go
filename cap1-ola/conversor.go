package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	// Testa a quantidade e argumentos passados por linha de comando
	if len(os.Args) < 3 {
		fmt.Println("Uso: conversor <valores> <unidades>")
		os.Exit(1)
	}

	// Declara a varivel como último argumento passado via linha de comando
	unidadeOrigem := os.Args[len(os.Args)-1]

	// Faz um slicing nos argumentos
	valoresOrigem := os.Args[1 : len(os.Args)-1]

	// Não podemos atribuir um valor a ela no momento da declaração
	var unidadeDestino string

	// Se unidadeOrigem for celsius define fahrenheit como unidadeDestino
	if unidadeOrigem == "celsius" {
		unidadeDestino := "fahrenheit"

	} else if unidadeOrigem == "quilometros" {
		// Se não for celsius e sim quilometros define unidadeDestino como milhas
		unidadeDestino := "milhas"

	} else {
		// Se não for nada disso Emite um erro
		fmt.Printf("Error: '%s' valor desconhecido\n", unidadeOrigem)
		os.Exit(1)

	}

	for i, v := range valoresOrigem {
		valorOrigem, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Printf(
				"O valor %s na posição %d não é um número válido!\n",
				v, i)
			os.Exit(1)
		}

		var valorDestino float64

		if unidadeOrigem == "celsius" {
			valorDestino = valorOrigem*1.8 + 32
		} else {
			valorDestino = valorOrigem / 1.60934
		}

		fmt.Printf("%.2f %s = %.2f %s\n", valorOrigem, unidadeOrigem, valorDestino, unidadeDestino)
		fmt.Printf("i=%d  -   v=%d \n", i, v)
	}

}
