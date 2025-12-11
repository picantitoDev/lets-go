package main

import "fmt"

func incrementar(n *int) {
	*n++
}

func main() {
	val := 10
	incrementar(&val)
	fmt.Println("Valor incrementado:", val)
}
