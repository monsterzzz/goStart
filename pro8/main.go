package main

import "fmt"

func main() {
	// 可变参数
	//changeArgs(1,2,3,4,5,6)
	fmt.Println(findMaxMin(2, 4, 6, 8))

}

func changeArgs(val ...int) {
	for i, v := range val {
		fmt.Printf("第 %d 个参数 -> %d\n", i+1, v)
	}
}

func sum(val ...int) int {
	var result int
	for _, v := range val {
		result += v
	}
	return result
}

func findMaxMin(val ...int) (int, int) {
	if len(val) == 0 {
		return 0, 0
	} else if len(val) == 1 {
		return val[0], val[0]
	}

	var maxV int
	var minV int

	maxV = val[0]
	minV = val[0]

	for _, v := range val {
		if v > maxV {
			maxV = v
		}
		if v < minV {
			minV = v
		}
	}
	return maxV, minV

}
