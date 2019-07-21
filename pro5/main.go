package main

import "fmt"

func main() {
	// 消除相邻重复字符串
	s := "aabbaaxxccasdaw"
	a := []rune(s)

	for i := 0; i < len(a); i++ {
		//fmt.Println(string(a))
		clear(&a[i], &a[i+1])

		if i == len(s)-2 {
			break
		}
	}

	b := make([]rune, 0)
	for _, v := range a {
		if v != 0 {
			b = append(b, v)
		}
	}
	fmt.Println(string(b))
	fmt.Println("111")

	//b := string(a)
	//b = strings.Replace(b," ","1",-1)

}

func clear(x, y *rune) {
	var n rune
	if *x == *y {
		*y = n
	}
}
