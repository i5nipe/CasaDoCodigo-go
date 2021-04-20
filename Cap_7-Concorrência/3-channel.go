package main

import "fmt"

func main() {
	// Criando um canal para trafegar informações do tipo
	// int.
	c := make(chan int)
	go produzir(c)

	// Recebendo um valor enviado do canal c
	valor := <-c
	fmt.Println(valor)
}

func produzir(c chan int) {
	// Enviando um valor para o canal c
	c <- 33
}
