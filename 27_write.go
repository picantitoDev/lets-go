package main

import "os"

func main() {
	os.WriteFile("log.txt", []byte("Casi termina el año"), 0644)
}
