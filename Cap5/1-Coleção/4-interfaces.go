// Uma interface é a definição de um conjunto de métodos comuns
// a um ou mais tipos. É o que permite a criação de tipos polimórficos.
package main

import "fmt"

type Operacao interface {
	Calcular() int
}

type Soma struct {
	operando1, operando2 int
}

func (s Soma) Calcular() int {
	return s.operando1 + s.operando2
}

func (s Soma) String() string {
	return fmt.Sprintf("%d + %d", s.operando1, s.operando2)
}

func main() {
	var soma Operacao
	soma = Soma{10, 20}
	fmt.Printf("%v = %d\n", soma, soma.Calcular())
}
