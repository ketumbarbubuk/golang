package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "selamat malam"
	for _, y := range input {
		fmt.Printf("%c\n", y)
	}
	res1 := strings.Count(input, "s")
	res2 := strings.Count(input, "e")
	res3 := strings.Count(input, "l")
	res4 := strings.Count(input, "a")
	res5 := strings.Count(input, "m")
	res6 := strings.Count(input, "t")
	res7 := strings.Count(input, " ")
	fmt.Printf("map [ : %d, a: %d, e: %d, l: %d, m: %d, s: %d, t: %d]", res7, res4, res2, res3, res5, res1, res6)
}
