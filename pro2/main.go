package main

import (
	. "awesomeProject2/pro2/mylib"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// 温度转换

	iptReader := bufio.NewScanner(os.Stdin)
	fmt.Println("please input a sheshidu num:")
	iptReader.Scan()

	num, err := strconv.Atoi(iptReader.Text())

	if err != nil {
		fmt.Println("error input!")
		os.Exit(1)
	}

	iptSheshidu := Sheshidu(num)

	var h1 Huashidu

	h1 = S2h(iptSheshidu)

	fmt.Println("fff:", h1)

}
