package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 编写一个程序wordfreq程序，报告输入文本中每个单词出现的频率。在第一次调用Scan前先调用input.Split(bufio.ScanWords)函数，这样可以按单词而不是按行输入。
	m := make(map[string]int)
	fmt.Println(m)

	iptReader := bufio.NewScanner(os.Stdin)
	//ipt_reader.Scan()
	iptReader.Split(bufio.ScanWords)
	for iptReader.Scan() {
		wordFreq(iptReader.Text(), m)
	}
	fmt.Println(m)
}

func wordFreq(s string, m map[string]int) {
	m[s]++
}
