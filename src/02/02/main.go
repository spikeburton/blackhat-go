package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 1; i <= 1024; i++ {
		s := fmt.Sprintf("scanme.nmap.org:%d", i)
		conn, err := net.Dial("tcp", s)
		if err != nil {
			continue
		}
		conn.Close()
		fmt.Println("Connection successful on port", i)
	}
}
