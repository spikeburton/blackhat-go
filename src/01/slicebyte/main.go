package main

import "fmt"

func main() {
	sb := make([]byte, 6)
	sb = append(sb, 0x40)
	fmt.Println(string(sb))
	fmt.Println(cap(sb))
	sb = append(sb, 0x46)
	sb = append(sb, 0x46)
	fmt.Println(string(sb))
	fmt.Println(cap(sb))
}
