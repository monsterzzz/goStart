package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type huashidu float32
type sheshidu float32

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

	iptSheshidu := sheshidu(num)

	var h1 huashidu

	h1 = s2h(iptSheshidu)

	fmt.Println("fff:", h1)

}

func s2h(s sheshidu) huashidu {
	var tmpS float32
	tmpS = float32(s)
	result := tmpS*1.8 + 32
	return huashidu(result)
}
