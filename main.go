package main

import "fmt"

func main() {
	var printValue string = "Hello World"
	printme(printValue)

	var numerator int = 11
	var denominator int = 2

	var result, remainder int = intDivision(numerator, denominator)
	fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
}

func printme(printvalue string) {
	fmt.Println(printvalue)
}

func intDivision(numerator int, denominator int) (int, int) {
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder
}
