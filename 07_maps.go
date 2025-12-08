package main

import "fmt"

func main() {
	precios := map[string]float64{
		"pan":   1.50,
		"leche": 0.99,
	}
	fmt.Println("El precio del pan es:", precios["pan"])
}
