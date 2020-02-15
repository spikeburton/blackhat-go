package main

import (
	"fmt"
	"os"
)

type person struct {
	fname string
	lname string
}

func main() {
	p := person{"Spike", "Burton"}
	fmt.Fprintf(os.Stdout, "Hello, %s %s!\n", p.fname, p.lname)
}
