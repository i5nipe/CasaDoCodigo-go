// Uma interface é a definição de um conjunto de métodos comuns
// a um ou mais tipos. É o que permite a criação de tipos polimórficos.
package main

import "fmt"

type Operacao interface {
	Calcular()
}

func main() {
	fmt.Println("vim-go")
}
