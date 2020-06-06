package main

import (
	"fmt"
	"net"
	"sort"
)

func worker(ports, results chan int) {
	for p := range ports {
		addr := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {
	ports := make(chan int, 1000)
	results := make(chan int)
	var open []int

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	for i := 0; i < 1024; i++ {
		port := <-results
		if port != 0 {
			open = append(open, port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(open)
	for _, port := range open {
		fmt.Printf("%d open\n", port)
	}
}
