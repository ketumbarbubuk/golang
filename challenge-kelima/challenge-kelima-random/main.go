package main

import (
	"fmt"
	"time"
)

func routine1(data interface{}) {
	for i := 1; i <= 4; i++ {
		fmt.Println("[bisa1 bisa2 bisa3] ", i, data)
		time.Sleep(500 * time.Millisecond)
	}
}

func routine2(data interface{}) {
	for i := 1; i <= 4; i++ {
		fmt.Println("[coba1 coba2 coba3] ", i, data)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	data1 := "1"
	data2 := 2

	for i := 1; i <= 2; i++ {
		go routine1(data1)
		go routine2(data2)
	}

	time.Sleep(5 * time.Second)
}
