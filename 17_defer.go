package main

import "fmt"

func main() {
	defer fmt.Println("Esto se ejecuta al final.")
	fmt.Println("Esto se ejecuta primero.")
}
