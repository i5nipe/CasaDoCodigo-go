// Uma struct é simplesmente uma coleção de
// variáveis que formam um novo tipo
package main

import "fmt"

type Arquivo struct {
	nome       string
	tamanho    float64
	caracteres int
	palavras   int
	linhas     int
}

func main() {
	// Podemos atribuir valores utilizando a mesma ordem em
	// que foram definidos:
	arquivo := Arquivo{"arquivo.txt", 12.68, 12986, 1862, 220}
	fmt.Println(arquivo) // {arquivo.txt 12.68 12986 1862 220}

	// Mas também é possível atribuir valores especificando
	// o nome e não é necessario definir todas as variaveis
	codigoFonte := Arquivo{tamanho: 1.12, nome: "programa.go"}
	fmt.Println(codigoFonte) // {programa.go 1.12 0 0 0}
}
