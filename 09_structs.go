package main

import "fmt"

type Persona struct {
	Nombre string
	Edad   int
}

func main() {
	p := Persona{Nombre: "Juan", Edad: 30}
	fmt.Printf("%s tiene %d años.\n", p.Nombre, p.Edad)
}
