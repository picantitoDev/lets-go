package main

import "fmt"

func Imprimir[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func main() {
	Imprimir([]int{1, 2, 3})
	Imprimir([]string{"a", "b"})
}
