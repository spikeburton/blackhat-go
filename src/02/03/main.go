package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 1; i <= 1024; i++ {
		go func(j int) {
			s := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", s)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Println("Connection successful on port", j)
		}(i)
	}
}
