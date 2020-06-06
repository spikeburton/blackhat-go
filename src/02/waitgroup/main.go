package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 65535; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			addr := fmt.Sprintf("127.0.0.1", j)
			conn, err := net.Dial("tcp", addr)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("Port %d open\n", j)
		}(i)
	}
	wg.Wait()
}
