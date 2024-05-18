package test

import (
	"fmt"
	"net/http"
)

func middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println("Antes do manipulador")
		h.ServeHTTP(w, r)
		println("Depois do manipulador")
	})
}

func main() {
	meuManipulador := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Olá, Mundo!")
	})

	http.Handle("/", middleware(meuManipulador))
	http.ListenAndServe(":8080", nil)
}
