package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		// 1 在这里需要你写算法
		// 2 要求每秒钟调用一次proc函数
		// 3 要求程序不能退出

		for {
			func() {
				defer func() {
					if err := recover(); err != nil { // revcover 只在defer的函数中有效
						fmt.Println(err)
					}
				}()
				proc()
			}()
			time.Sleep(time.Second)
			//defer func() {
			//	if err := recover(); err != nil { //error 直接退出当前goroutine
			//		fmt.Println(err)
			//	}
			//}()
			//proc()
			//time.Sleep(time.Second)
		}

	}()

	select {}
}

func proc() {
	panic("ok")
}
