package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, _ := http.Get("https://google.com")
	fmt.Println("Status:", resp.Status)
}
