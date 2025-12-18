package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() { time.Sleep(time.Second); c1 <- "uno" }()
	go func() { time.Sleep(time.Second); c2 <- "dos" }()
	for i := 0; i < 2; i++ {
		select {
		case m1 := <-c1:
			fmt.Println(m1)
		case m2 := <-c2:
			fmt.Println(m2)
		}
	}
}
