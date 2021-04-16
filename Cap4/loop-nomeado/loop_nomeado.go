package main

import "fmt"

func main() {
	var i int

	//Estamos nomeando o loop para poder sair dele de dentro do switch:
loop:
	for i = 0; i < 10; i++ {
		// Se i for ...
		switch i {
		// 2 ou 3 faça:
		case 2, 3:
			fmt.Printf("Quebrando switch, i == %d\n", i)
			//Usando o break sem argumentos para sair do switch
			break
		// 5 faça:
		case 5:
			fmt.Println("Quebrando loop, i == 5.")
			//Usando o break com o nome do bloco como argumento nós saimos do for
			break loop
		}
	}
	fmt.Println("Fim")
}
