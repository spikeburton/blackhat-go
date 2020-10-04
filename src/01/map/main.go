package main

import (
	"fmt"
)

func main() {
	var m = make(map[string]string)
	m["some key"] = "hello world"
	fmt.Println(m)
}
