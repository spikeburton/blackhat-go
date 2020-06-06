package main

import "fmt"

// Person is a type
type Person struct {
	Name string
	Age  int
}

func (p *Person) sayHello() {
	fmt.Printf("Hello, %s\n", p.Name)
}

func main() {
	var guy = new(Person)
	guy.Name = "Dave"
	guy.sayHello()
}
