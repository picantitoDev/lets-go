package main

import "fmt"

type Saludador interface{ Saludar() string }
type Persona struct{ Nombre string }

func (p Persona) Saludar() string { return "Hola, soy " + p.Nombre }

func main() {
	var s Saludador = Persona{"Gopher"}
	fmt.Println(s.Saludar())
}
