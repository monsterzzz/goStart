package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "12,34,56,"
	fmt.Println(s[:strings.LastIndexAny(s[:len(s)-1], ",")])

}
