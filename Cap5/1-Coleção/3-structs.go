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
	arquivo := Arquivo{"arquivo.txt", 12.68, 12986, 1862, 220}
	fmt.Println(arquivo)
}
