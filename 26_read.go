package main

import (
	"fmt"
	"os"
)

func main() {
	dat, _ := os.ReadFile("26_read.go")
	fmt.Print(string(dat))
}
