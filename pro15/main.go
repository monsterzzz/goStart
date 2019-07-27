package main

import (
	"fmt"
	"math/rand"
	"time"
)

type celsiusFlag struct {
	name string
}

func main() {
	//for i := 0 ; i< 10;i++{
	//	fmt.Println(rand.Intn(30))
	//}

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			randInt := rand.Intn(3)
			time.Sleep(time.Duration(randInt) * time.Second)
			s := fmt.Sprintf("I am c1,I have sleep %d s", randInt)
			c1 <- s
		}
	}()

	go func() {
		for {
			randInt := rand.Intn(3)
			time.Sleep(time.Duration(randInt) * time.Second)
			s := fmt.Sprintf("I am c2,I have sleep %d s", randInt)
			c2 <- s
		}
	}()

	for {
		select {
		case <-c1:
			fmt.Printf("c1 -> %s\n", <-c1)

		case <-c2:
			fmt.Printf("c2 -> %s\n", <-c2)
		}
	}
}
