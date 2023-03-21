package main

import (
	"fmt"
	"sync"
	"time"
)

func routine1(data interface{}, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()

	for i := 1; i <= 4; i++ {
		mu.Lock()
		fmt.Println("[bisa1 bisa2 bisa3] ", i)
		time.Sleep(500 * time.Millisecond)
		mu.Unlock()
	}
}

func routine2(data interface{}, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()

	for i := 1; i <= 4; i++ {
		mu.Lock()
		fmt.Println("[coba1 coba2 coba3] ", i)
		time.Sleep(500 * time.Millisecond)
		mu.Unlock()
	}
}

func main() {
	data1 := "1"
	data2 := 2

	var wg sync.WaitGroup
	wg.Add(2)

	var mu sync.Mutex

	go routine2(data2, &wg, &mu)
	go routine1(data1, &wg, &mu)

	wg.Wait()
}
