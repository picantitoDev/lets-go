package main

import "fmt"

func main() {
	dia := "Lunes"
	switch dia {
	case "Lunes":
		fmt.Println("Inicio de semana")
	default:
		fmt.Println("Otro día")
	}
}
