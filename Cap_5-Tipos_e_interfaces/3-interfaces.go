// Uma interface é a definição de um conjunto de métodos comuns
// a um ou mais tipos. É o que permite a criação de tipos polimórficos.
package main

import "fmt"

//Definindo tipo operacão como interface
type Operacao interface {
	// Método Calcular não aceita argumentos e retorna int.
	Calcular() int
}

// Definindo tipo Soma
type Soma struct {
	operando1, operando2 int
}

type Subtracao struct {
	operando1, operando2 int
}

// Metodo para calcular os operandos
func (s Soma) Calcular() int {
	return s.operando1 + s.operando2
}

func (s Soma) String() string {
	return fmt.Sprintf("%d + %d", s.operando1, s.operando2)
}

func (s Subtracao) Calcular() int {
	return s.operando1 + s.operando2
}

func (s Subtracao) String() string {
	return fmt.Sprintf("%d - %d", s.operando1, s.operando2)
}

func main() {
	// Criando um slice de Operacao e atribuindo valores
	operacoes := make([]Operacao, 4)
	operacoes[0] = Soma{10, 20}
	operacoes[1] = Subtracao{30, 15}
	operacoes[2] = Subtracao{10, 50}
	operacoes[3] = Soma{5, 2}

	acumulador := 0
	for _, op := range operacoes {
		valor := op.Calcular()
		fmt.Printf("%v = %d\n", op, valor)
		acumulador += valor
	}
	fmt.Println("Valor acumulado =", acumulador)
}
