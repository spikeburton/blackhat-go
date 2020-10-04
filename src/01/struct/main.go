package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p *Person) SayHello() {
	fmt.Printf("Hello, %s\n", p.Name)
}

func main() {
	var guy = new(Person)
	guy.Name = "Dave"
	guy.SayHello()
}
