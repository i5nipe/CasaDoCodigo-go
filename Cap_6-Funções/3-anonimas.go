package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	// o \b é para especificar todos os numerais
	// e como esta dentro de um string precisa ser escapado \\b
	expr := regexp.MustCompile("\\b\\w")

	// Aqui está a tal função anônima, ela foi definida
	// dentro de uma variavel e não possui um nome
	transformadora := func(s string) string {
		return strings.ToUpper(s)
	}
	texto := "antonio carlos macaco banana tokio intabivel"

	fmt.Println(transformadora(texto))
	fmt.Println(expr.ReplaceAllStringFunc(texto, transformadora))
}
