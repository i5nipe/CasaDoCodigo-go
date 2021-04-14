package main

import (
	"fmt"
	"os"
	"strings"
)

// Conta a frequência de palavras com as mesmas iniciais.
func main() {
	palavras := os.Args[1:]

	testa := colherEstatisticas(palavras)

	imprimir(testa)
}

// map, array associativo ou dicionário tem funcionamento
// parecido com os dicionários de python
// é uma coleção de pares chave/valor em que cada chave é unica
// e armazena um único valor.
// Nessa função a chave é a letra inicial da palavra e o valor
// é a quantidade de palavras com esta inicial.
func colherEstatisticas(palavras []string) map[string]int {
	estatisticas := make(map[string]int)

	for _, palavra := range palavras {

		// Converte a inicial da palavra para maiuscula e armazena na variavel.
		// Se retirado os dois pontos ":" a função string retorna um erro
		// reclamando da entrada do tipo byte.
		inicial := strings.ToUpper(palavra[:1])

		// Estatisticas chama o map da inicial e retorna dois dados
		// o primeiro é o valor daquela chave, no nosso caso um int
		// o segundo é um bool que volta false se não encontrar a chave
		// e true se encontrar
		contador, encontrado := estatisticas[inicial]

		if encontrado {
			estatisticas[inicial] = contador + 1
		} else {
			estatisticas[inicial] = 1
		}

	}
	return estatisticas
}

func imprimir(estatisticas map[string]int) {
	fmt.Println("Contagem de palavras iniciadas em cada letra:")

	// Quando usado em um map, o range retorna a chave e o valor de map
	for inicial, contador := range estatisticas {
		fmt.Printf("%s = %d\n", inicial, contador)
	}
}
