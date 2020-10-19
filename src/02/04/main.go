package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			s := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", s)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Println("Connection successful on port", j)
		}(i)
	}
	wg.Wait()
}
