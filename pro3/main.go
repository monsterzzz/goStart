package main

import "fmt"

func main() {
	// 反转silce

	a := "abcwa"
	//s1 := [len(s)] string{}
	s := []rune(a)

	for a := 0; a < len(s); a++ {
		if len(s) == 0 && len(s) == 1 {
			break
		}

		swap(&s[a], &s[len(s)-a-1])

		if len(s)%2 == 0 {
			if a == len(s)/2-1 {
				break
			}
		} else {
			if a == (len(s) / 2) {
				fmt.Println("!")
				break
			}
		}
	}
	fmt.Println(string(s))
}

func swap(x, y *rune) {
	*x, *y = *y, *x
}
