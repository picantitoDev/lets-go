package main

import "fmt"

func main() {
	frutas := []string{"manzana", "banana"}
	frutas = append(frutas, "cereza")
	fmt.Println("Lista de frutas:", frutas)
}
