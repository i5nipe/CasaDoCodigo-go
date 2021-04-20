package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc(
		// Especificamos que o http.HandleFunc atenderá as requisições
		// da URL /tempo .
		"/tempo",
		func(w http.ResponseWriter, r *http.Request) {
			// Utilizamos a função Fprintf pela primeira vez,
			// Ela funciona  semenhante a Printf com um porém
			// Seu primeiro argumento deve receber um io.Writer
			// onde sua saida sera escrita.
			fmt.Fprintf(w, "%s",
				time.Now().Format("2006-01-02 15:04:05"))
		})

	http.ListenAndServe(":8080", nil)
}
