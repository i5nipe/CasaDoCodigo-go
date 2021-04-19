package main

import "fmt"

// Criando um tipo
type Estado struct {
	nome      string
	populacao int
	capital   string
}

func main() {
	estados := make(map[string]Estado, 3)

	estados["GO"] = Estado{"Goiás", 6434052, "Goiania"}
	estados["PB"] = Estado{"Paraíba", 3914418, "João Pessoa"}
	estados["AM"] = Estado{"Amazonas", 3807923, "Manaus"}

	fmt.Println(estados) // map[AM:{Amazonas 3807923 Manaus} GO:{Goiás 6434052 ... ]
}
