// Uma struct é simplesmente uma coleção de
// variáveis que formam um novo tipo
package main

import "fmt"

type Arquivo struct {
	nome       string
	tamanho    float64
	caracteres int
	palavras   int
	linhas     int
}

func main() {
	// ----------------------{ Definindo }-------------------------
	// Podemos atribuir valores utilizando a mesma ordem em
	// que foram definidos:
	arquivo := Arquivo{"arquivo.txt", 12.68, 12986, 1862, 220}
	fmt.Println(arquivo) // {arquivo.txt 12.68 12986 1862 220}

	// Mas também é possível atribuir valores especificando
	// o nome e não é necessario definir todas as variaveis
	codigoFonte := Arquivo{tamanho: 1.12, nome: "programa.go"}
	fmt.Println(codigoFonte) // {programa.go 1.12 0 0 0}

	// ----------------------{ Acessando }-------------------------
	// Podemos utilizar o nomedavariavel.campo como no exemplo abaixo:
	fmt.Printf("%s\t%.2fKB\n", arquivo.nome, arquivo.tamanho)
	fmt.Printf("%s\t%.2fKB\n", codigoFonte.nome, codigoFonte.tamanho)

	// ----------------------{ Modificando }-------------------------
	// Qualquer valor armazenado em uma struct pode ser alterado,
	// eles são tipos mutáveis
	fmt.Println(codigoFonte) // {programa.go 1.12 0 0 0}
	codigoFonte.nome = "Modificado.go"
	fmt.Println(codigoFonte) // {Modificado.go 1.12 0 0 0}

	// ----------------------{ Métodos }-------------------------
	// Podemos acessar campos dos nossos tipos customizados
	// em métodos.
	arquivo = Arquivo{"arquivo.txt", 12.68, 12986, 1862, 220}
	fmt.Printf("Media de palavras por linha: %2.f\n",
		arquivo.MediaDePalavrasPorLinha())
	fmt.Printf("Tamanho médio de palavras: %.2f",
		arquivo.TamanhoMedioDePalavras())
}

func (arq *Arquivo) TamanhoMedioDePalavras() float64 {
	// Estamos convertendo inteiro para float pois não
	// queremos que parte decimal da divisão seja truncada
	return float64(arq.caracteres) / float64(arq.palavras)
}

func (arq *Arquivo) MediaDePalavrasPorLinha() float64 {
	return float64(arq.palavras) / float64(arq.linhas)
}
