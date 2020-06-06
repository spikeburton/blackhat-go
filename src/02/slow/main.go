package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 1; i <= 1024; i++ {
		addr := fmt.Sprintf("scanme.nmap.org:%d", i)
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			fmt.Printf("i is %d\n", i)
			continue
		}
		conn.Close()
		fmt.Printf("Connected to port %d\n", i)
	}
}
