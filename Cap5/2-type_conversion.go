package main

import "fmt"

type ListaDeCompras []string

func imprimirSlice(slice []string) {
	fmt.Println("Slice: ", slice)
}

func imprimirLista(lista ListaDeCompras) {
	fmt.Println("Lista: ", lista)
}

func main() {
	lista := ListaDeCompras{"Alface", "Atum", "Leite"}
	slice := []string{"Alface", "Atum", "Leite"}

	imprimirSlice([]string(lista))
	imprimirLista(ListaDeCompras(slice))
}
