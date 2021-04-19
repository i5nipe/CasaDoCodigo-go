// Um map é uma coleção de pares chave-valor sem ordem definida. Também conhecida
// como hashtable, dicionário de dados , array associativo, entre outros nomes.
package main

import "fmt"

func main() {
	// Qualquer tipo que suporte os operadores de igualdade(== !=) pode
	// ser usado como chaves de maps.
	// As chaves do map devem ser do mesmo tipo, os valores também!
	// Porém os valores aceitam o tipo interface{} com isso podemos fazem
	// um type assertion quando recuperar os valores.

	// Declarando maps vazios, os dois são equivalentes.
	// Com chaves do tipo int e valores do tipo string.
	vazio1 := map[int]string{}
	vazio2 := make(map[int]string)
	fmt.Println(vazio1, vazio2)

	// A quantidade de valores armazenados em um map é flexível e pode crescer.
	// Mas podemos especificar sua capacidade com um segundo argumento no make.
	// Fazer isso é altamente importante pois tornamos o usa de memória eficiente
	mapaGrande := make(map[int]string, 4096)
	fmt.Println(mapaGrande)

	// Declarando e adicionando valores em um map
	capitais := map[string]string{
		"GO": "Goiania",
		"PB": "João Pessoa",
		"PR": "Curitiba"}
	// Adicionando valores em um map
	capitais["AM"] = "Manaus"
	capitais["SE"] = "Aracaju"

	// Podemos verificar o tamanho de maps com len
	fmt.Println(len(capitais)) // 5

	populacao := make(map[string]int, 6)
	populacao["GO"] = 6434052
	populacao["PB"] = 3914418

	fmt.Println(capitais)  // map[GO:Goiania PB:João Pessoa ...]
	fmt.Println(populacao) // map[GO:6434052 PB:3914418]

}
