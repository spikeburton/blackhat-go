package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 0; i <= 1024; i++ {
		go func(j int) {
			addr := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", addr)
			if err != nil {
				// fmt.Printf("Failed to connect to port %d\n", i)
				return
			}
			conn.Close()
			fmt.Printf("Connected to port %d\n", j)
		}(i)
	}
}
