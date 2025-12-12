package main

import "fmt"

type Rectangulo struct {
	Ancho, Alto float64
}

func (r Rectangulo) Area() float64 {
	return r.Ancho * r.Alto
}

func main() {
	rect := Rectangulo{Ancho: 10, Alto: 5}
	fmt.Println("Área del rectángulo:", rect.Area())
}
