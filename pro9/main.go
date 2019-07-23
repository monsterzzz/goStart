package main

import "fmt"

func main() {
	// 可变参数
	//defer panic recover

	// defer

	//a := 5
	//fmt.Println(a)
	//defer func() {fmt.Println(a)}()
	//a++

	//panic
	// 个人理解是类似主动抛出未预料的异常

	// recover 能够在defer中捕捉接下来的异常
	//defer func() {	fmt.Println(recover())}()
	//myPanic()

	defer func() { fmt.Println(recover()) }()
	divNum(0, 0)

}

func myPanic() {
	panic(5)
}

func divNum(a, b int) int {
	return a / b
}
