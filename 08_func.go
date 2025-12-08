package main

import "fmt"

func swap(a, b string) (string, string) {
	return b, a
}

func main() {
	primero, segundo := swap("hola", "mundo")
	fmt.Println(primero, segundo)
}
