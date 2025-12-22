package main

import (
	"encoding/json"
	"fmt"
)

type Post struct {
	Title string `json:"titulo"`
}

func main() {
	p := Post{Title: "Aprendiendo Go"}
	b, _ := json.Marshal(p)
	fmt.Println(string(b))
}
