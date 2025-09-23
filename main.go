package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, serveur Go fonctionne ! 🚀")
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erreur serveur:", err)
	}
}
