package main

import (
	"fmt"
	"os"
)

// Uma variadic_function pode receber zero ou mais argumentos
// do tipo especificado.
func CriarArquivos(dirBase string, arquivos ...string) {
	for _, nome := range arquivos {
		caminhoArquivo := fmt.Sprintf(
			"%s/%s.%s", dirBase, nome, "txt")

		arq, err := os.Create(caminhoArquivo)

		// defer é usado para instruir o Compilador a executar
		// imediatamente antes da função atual retornar.
		defer arq.Close()

		if err != nil {
			fmt.Printf("Erro ao criar arquivo %s: %v\n",
				nome, err)
			os.Exit(1)
		}
		fmt.Printf("Arquivo %s criado.\n", arq.Name())
	}
}
func main() {
	fmt.Println("vim-go")
}
