package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Feliz Navidad desde un servidor Go")
	})
	fmt.Println("Servidor listo en :8080")
}
