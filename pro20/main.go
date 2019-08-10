package main

import (
	"fmt"
	"time"
)

func main() {
	s1 := "a"
	s2 := "b"

	c1 := make(chan string, 3)
	c2 := make(chan string)
	go func() {
		n := 0
		for {
			c1 <- fmt.Sprintf("%s%d", s1, n)
			n++
			time.Sleep(1 * time.Second)
			if n == 10 {
				close(c1)
				break
			}

		}
	}()
	go func() {
		n1 := 0
		for {
			c2 <- fmt.Sprintf("%s%d", s2, n1)
			n1++
			time.Sleep(1 * time.Second)
			if n1 == 10 {
				close(c2)
				break
			}

		}
	}()

	//for i := range c1{
	//	fmt.Println(i)
	//	//fmt.Println(<-c2)
	//	//select {
	//	//case <-c1:
	//	//	fmt.Println(<-c1)
	//	//case <-c2:
	//	//	fmt.Println(<-c2)
	//	//}
	//}

	for {
		select {
		case <-c1:
			fmt.Println(<-c1)
		case <-c2:
			fmt.Println(<-c2)
		}
	}

}
