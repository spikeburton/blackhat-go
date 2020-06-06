package main

import "fmt"

// Person type
type Person struct {
	Name string
	Age  int
}

// Dog type
type Dog struct {}

func (d *Dog) sayHello() {
	fmt.Println("Woof woof")
}

// Friend type
type Friend interface {
	sayHello()
}

func (p *Person) sayHello() {
	fmt.Printf("Hello, %s\n", p.Name)
}

// Greet func
func Greet(f Friend) {
	f.sayHello()
}

func main() {
	var guy = new(Person)
	guy.Name = "Spike"
	Greet(guy)
	var dog = new(Dog)
	Greet(dog)
}
