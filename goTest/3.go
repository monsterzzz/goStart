package main

import (
	"fmt"
)

func main() {
	var x string = nil // string can not be nil,nil must be a *
	if x == nil {
		x = "default"
	}
	fmt.Println(x)
}
