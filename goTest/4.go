package main

import (
	"fmt"
)

const (
	a = iota // 0
	b = iota // 1
)

// re count
const (
	name = "menglu" //iota = 0
	c    = iota     // 1
	d    = iota     // 2
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
