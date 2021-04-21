package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	// Testa a quantidade de argumentos passados por linha de comando
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
		unidadeDestino = "fahrenheit"

	} else if unidadeOrigem == "quilometros" {
		// Se não for celsius e sim quilometros define unidadeDestino como milhas
		unidadeDestino = "milhas"

	} else if unidadeOrigem == "metros" {
		unidadeDestino = "feet"
	} else {
		// Se não for nada disso Emite um erro
		fmt.Printf("Error: '%s' valor desconhecido\n", unidadeOrigem)
		os.Exit(1)

	}

	// range retorna dois valores um o index e o outro o valor em si
	for i, v := range valoresOrigem {
		// ParseFloat retorna dois valores, o valor da conversão e um valor de error
		valorOrigem, err := strconv.ParseFloat(v, 64)

		// Se err retornar nil é porque algo ocoreu errado
		if err != nil {
			fmt.Printf(
				"O valor %s na posição %d "+
					"não é um número válido!\n",
				v, i)
			os.Exit(1)
		}

		var valorDestino float64

		if unidadeOrigem == "celsius" {
			valorDestino = valorOrigem*1.8 + 32
		} else if unidadeOrigem == "quilometros" {
			valorDestino = valorOrigem / 1.60934
		} else {
			valorDestino = valorOrigem * 3.281
		}

		// Imprime os resultados
		fmt.Printf("%.2f %s = %.2f %s\n",
			valorOrigem, unidadeOrigem, valorDestino, unidadeDestino)
	}

}
