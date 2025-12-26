package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	v := 0
	mu.Lock()
	v++
	mu.Unlock()
	fmt.Println(v)
}
