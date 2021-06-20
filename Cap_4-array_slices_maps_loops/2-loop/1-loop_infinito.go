package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//Precisamos de uma seed para geração de números "aleatorios"
	// Nesse casa estamos usando o timestamp do momento da execução "time.Now()"
	// E o número de nanosegundos desde 1° de janeiro de 1970 "UnixNano()"
	rand.Seed(time.Now().UnixNano())
	n := 0

	//Inicia um loop infinito
	for {
		n++

		// A função Intn gera um número pseudo aleatório entre 0 até 4200
		i := rand.Intn(4200)
		fmt.Println(i)

		// Se o resto da divição de i por 42 for igual a zero
		if i%42 == 0 {
			// Por padrão o Break sai do loop mais próximo
			break
		}
	}

	fmt.Printf("Saída após %v iterações.\n", n)
}
