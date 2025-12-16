package main

import (
	"fmt"
	"time"
)

func mensaje(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
	}
}

func main() {
	go mensaje("Corriendo en segundo plano...")
	mensaje("Corriendo en hilo principal")
	time.Sleep(time.Millisecond * 100)
}
