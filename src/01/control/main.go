package main

import "fmt"

func main() {
	nums := []int{2, 4, 6, 8}
	for idx, val := range nums {
		fmt.Printf("num %d is %d\n", idx, val)
	}
	ppl := map[string]int{"Spike": 33, "David": 27, "Peter": 28}
	for k, v := range ppl {
		fmt.Printf("%s is %d years old\n", k, v)
	}
}
