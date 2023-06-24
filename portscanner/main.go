// In the book, the coding does not take care about concurreny problem! We need to use WaitGroup for fixing it.
package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	fmt.Println("Start Scanning")
	start := time.Now()

	var wg sync.WaitGroup
	wg.Add(1024)
	for i := 1; i <= 1024; i++ {
		go func(j int) {
			address := fmt.Sprintf("192.168.0.61:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				wg.Done()
				return
			}
			conn.Close()
			fmt.Printf("%d open\n", j)
			wg.Done()
		}(i)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println("Scanning duration ", elapsed)
}
