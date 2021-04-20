package main

import "fmt"

func main() {
	// Criamos um channel com buffer de 3
	c := make(chan int, 3)
	go produzir(c)

	// Definimos um buffer de 3 e estamos enviando 4 valores
	// iriamos recebe no terminal um erro deadlock se não
	// fosse o close no final da função produzir()
	for valor := range c {
		fmt.Println(valor)
	}
}

func produzir(c chan int) {
	c <- 1
	c <- 2
	c <- 3

	close(c)
}
