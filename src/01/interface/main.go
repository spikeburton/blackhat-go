package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

type Dog struct{}

type Friend interface {
	SayHello()
}

func (p *Person) SayHello() {
	fmt.Printf("Hello, %s\n", p.Name)
}

func (d *Dog) SayHello() {
	fmt.Println("Woof woof")
}

func Greet(f Friend) {
	f.SayHello()
}

func main() {
	var guy = new(Person)
	guy.Name = "Dave"
	Greet(guy)
	var dog = new(Dog)
	Greet(dog)
}
