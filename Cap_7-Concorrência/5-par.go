package main

import "fmt"

func separar(nums []int, i, p chan<- int, pronto chan<- bool) {
	for _, n := range nums {
		if n%2 == 0 {
			p <- n
		} else {
			i <- n
		}
	}
	pronto <- true
}
func main() {
	i, p := make(chan int), make(chan int)
	pronto := make(chan bool)
	nums := []int{1, 2, 3, 42, 2, 324, 234, 2354, 2, 2, 112, 1, 123, 34, 345, 3, 2, 5, 346, 346, 3, 32222}
	go separar(nums, i, p, pronto)

	var impares, pares []int
	fim := false

	for !fim {
		select {
		case n := <-i:
			impares = append(impares, n)
		case n := <-p:
			pares = append(pares, n)
		case fim = <-pronto:
		}
	}
	fmt.Printf("Ímpares: %v | Pares: %v\n", impares, pares)
}
