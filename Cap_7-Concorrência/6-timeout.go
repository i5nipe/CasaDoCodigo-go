package main

import (
	"fmt"
	"time"
)

func executar(c chan<- bool) {
	time.Sleep(5 * time.Second)
	c <- true
}
func main() {
	m := make(chan bool, 1)
	go executar(c)

	fmt.Println("Esperando...")
}
