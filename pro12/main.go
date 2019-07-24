package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"unicode"
)

func main() {
	s := "hello，世界"
	for _, v := range s {
		fmt.Println(v)
	}
	var b bytes.Buffer
	b.WriteByte('s')
	fmt.Println(b)

	//var a = "12s34567890"
	var a = "123456789.0"
	i, e := addS(a)
	if e != nil {
		log.Fatal("fail adds")
	}
	fmt.Println(i)
}

func addS(s string) (string, error) {
	var b bytes.Buffer
	var e error

	n := 0
	headNum := len(s) % 3

	for i, v := range s {
		if !unicode.IsDigit(v) {
			e = errors.New("error string")
			break
		}
		if i == headNum {
			b.WriteByte(',')
		} else {
			if n == 3 {
				b.WriteByte(',')
				n = 0
			}
			n++
		}
		b.WriteByte(byte(v))
	}
	return b.String(), e
}
