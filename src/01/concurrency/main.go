package main

import (
	"fmt"
)

func f() {
	fmt.Println("Function f")
}

func strlen(s string, c chan int) {
	c <- len(s)
}

func main() {
	// go f()
	// time.Sleep(1 * time.Second)
	// fmt.Println("main function")
	c := make(chan int)
	go strlen("Salutations", c)
	go strlen("World", c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}
