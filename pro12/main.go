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
	var a = "1234567890"
	i, e := addS(a)
	if e != nil {
		log.Fatal("fail adds")
	}
	fmt.Println(i)

	fmt.Println(compareS("ba", "abc"))
	fmt.Println(sort("ba"))

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

func compareS(s1, s2 string) bool {
	var b bool
	b = true
	if s1 == s2 {
		b = false
		return b
	}
	if len(s1) != len(s2) {
		return false
	}
	if sort(s1) != sort(s2) {
		return false
	}
	return true
}

func sort(s string) string {
	b := []rune(s)
	for i := 0; i < len(b)-1; i++ {
		for j := i + 1; j < len(b); j++ {
			if b[i] > b[j] {
				b[i], b[j] = b[j], b[i]
			}
		}
	}
	return string(b)
}
