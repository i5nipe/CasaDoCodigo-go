package main

import "fmt"

type ListaDeCompras []string

// Uma função que precisa de uma string para funcionar
func imprimirSlice(slice []string) {
	fmt.Println("Slice: ", slice)
}

// Uma função que precisa de uma ListaDeCompras para funcionar
func imprimirLista(lista ListaDeCompras) {
	fmt.Println("Lista: ", lista)
}

func main() {
	lista := ListaDeCompras{"Alface", "Atum", "Leite"}
	slice := []string{"Alface", "Atum", "Leite"}

	// Se reparar bem nós estamos enviando a variavel errada
	// para as função, de acordo com seu tipo, Porem
	// o type conversion aqui está funcionando, Pois são
	// tipos compativeis.
	imprimirSlice([]string(lista))
	imprimirLista(ListaDeCompras(slice))
}
