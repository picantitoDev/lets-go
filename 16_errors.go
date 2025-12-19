package main

import (
	"errors"
	"fmt"
)

func validar(n int) (int, error) {
	if n < 0 {
		return -1, errors.New("no puede ser negativo")
	}
	return n * 2, nil
}

func main() {
	if _, err := validar(-1); err != nil {
		fmt.Println("Error:", err)
	}
}
