package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := "练习 4.8： 修改charcount程序，使用unicode.IsLetter等相关的函数，统计字母、数字等Unicode中不同的字符类别。"
	m := make(map[string]int)
	charCount(s, m)
	fmt.Println(m)

}

func charCount(s string, m map[string]int) {
	r := []rune(s)
	for _, v := range r {
		if unicode.IsDigit(v) {
			m["num"]++
		} else if unicode.IsLetter(v) {
			m["letter"]++
		} else {
			m["other"]++
			fmt.Println(v, string(v))
		}
		//fmt.Println(v,string(v),unicode.IsDigit(v))
	}
}
