package main

import (
	"fmt"
	"net"
	"sync"
)

// TCPscan - однопоточный сканер TCP портов
func tcpScan() {
	var wg sync.WaitGroup
	for i := 1; i < 1024; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			addr := fmt.Sprintf("scanme.nmap.org:%d", j)
			connect, err := net.Dial("tcp", addr)
			if err != nil {
				// порт закрыт или отфильтрован брендмауэром
				fmt.Printf("Порт %d закрыт или отфильтрован\n", j)
				return
			}
			connect.Close()
			fmt.Printf("Порт %v открыт\n", connect)
		}(i)
	}
	wg.Wait()
}
