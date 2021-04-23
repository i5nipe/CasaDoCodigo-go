package main

import "fmt"

// Uma função basica que não recebe nenhum argumento
// e não retorna nada.
func ImprimirVersao() {
	fmt.Println("0.666")
}

// Função que recebe um nome e idade e não retorna nada.
// Similar a declaração de variaveis a declaração de argumentos
// recebe um nome primeiro logo depois seu tipo e separando por ,
func ImprimirDados(nome string, idade int) {
	fmt.Printf("%s, %d anos.\n", nome, idade)
}

// Se o tipo do argumento for o mesmo podemos omiti-lo.
func ImprimirSoma(a, b int, texto string) {
	fmt.Printf("%s: %d\n", texto, a+b)
}

// Podemos retornar multiplos valores com o return
func PrecoFinal(precoCustoDolar float64) (float64, float64) {
	fatorLucro := 1.66
	taxaConversao := 2.85

	precoFinalDolar := precoCustoDolar * fatorLucro

	return precoFinalDolar, precoCustoDolar * taxaConversao
}

// Se não especificar uma saida para o return ele vai pegar as
// duas variaveis que definimos ser a saida da função e retorna elas.
func ConversaoFinal(precoCusto float64) (precoDolar, precoReal float64) {
	fatorLucro := 1.50
	taxaConversao := 3.00

	// Apenas atribuimos valor as variaveis
	// pois elas já foram definidas na declaração da função.
	precoDolar = precoCusto * fatorLucro
	precoReal = precoDolar * taxaConversao

	// Não é recomendado utilizar esses retornos pois
	// dificulta a leitura do codigo.
	return
}
func main() {
	ImprimirVersao()
	ImprimirDados("Matheus", 666)
	ImprimirSoma(5, 5, "Soma ae seu bosta")
	fmt.Println(PrecoFinal(5.62))

	precoDolar, precoReal := ConversaoFinal(65.8)
	fmt.Printf("Preço final em dólar: %.2f\n"+
		"Preço final em reais: %.2f\n", precoDolar, precoReal)
}
