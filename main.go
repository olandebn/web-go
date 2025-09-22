package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Route de test
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, serveur Go fonctionne ! 🚀")
	})

	// Lance le serveur sur le port 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erreur serveur:", err)
	}
}
