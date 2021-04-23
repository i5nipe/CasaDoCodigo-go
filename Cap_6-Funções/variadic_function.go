package main

import (
	"fmt"
	"os"
)

func CriarArquivos(dirBase string, arquivos ...string) {
	for _, nome := range arquivos {
		caminhoArquivo := fmt.Sprintf(
			"%s/%s.%s", dirBase, nome, "txt")

		arq, err := os.Create(caminhoArquivo)

		defer arq.Close()
	}
}
func main() {
	fmt.Println("vim-go")
}
