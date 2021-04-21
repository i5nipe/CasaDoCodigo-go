package main

import "fmt"

// Estamos criando praticamente um alias
// alias ListaDeCompras="string"
// Isso pode ser usado para melhorar a
// legibilidade do programa.
type ListaDeCompras []string

func main() {
	lista := make(ListaDeCompras, 3)
	lista[0] = "Alface"
	lista[1] = "Contra file"
	lista[2] = "Azeite"

	fmt.Println(lista) // [Alface Pepino Azeite]

	vegetais, carnes, sla := lista.Categorizar()
	fmt.Println("Vegetais: ", vegetais)
	fmt.Println("Carnes: ", carnes)
	fmt.Println("Outros: ", sla)
}

// A grande vantagem em tipos customizados é que
// podemos estendê-los, um exemplo é esse método abaixo,
// ele estamos recebendo um valor do tipo ListaDeCompras e
// expandindo para três strings.
func (lista ListaDeCompras) Categorizar() (
	[]string, []string, []string) {

	var vegetais, carnes, outros []string

	for _, line := range lista {
		switch line {
		case "Alface":
			vegetais = append(vegetais, line)
		case "Contra file":
			carnes = append(carnes, line)
		default:
			outros = append(outros, line)
		}
	}
	return vegetais, carnes, outros
}
