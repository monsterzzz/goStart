package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

type myWriter struct {
	data []byte
	w    io.Writer
}

type oldWriter struct {
	data []byte
}

func check(e error) {
	if e != nil {
		log.Fatal("error!")
	}
}
func main() {

	f, err := os.Open("./pro13/a.txt")
	if err != nil {
		log.Fatal("file not found!")
	}

	fmt.Println(bytesCounter(f))

}

func bytesCounter(f *os.File) (int, int) {
	data, err := ioutil.ReadAll(f)
	check(err)
	var wordCount, lineCount int
	var bLi []byte
	for i, v := range data {
		if v == 10 || i == len(data)-1 {
			lineCount++
			initCount := 1
			for _, v2 := range bLi {
				if v2 == 32 {
					initCount++
				}
			}
			wordCount += initCount
			bLi = []byte("")
		}
		bLi = append(bLi, v)
	}
	return wordCount, lineCount
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	s := "hello,world"
	b := []byte(s)
	var lenB int64
	for range b {
		lenB++
	}
	var mW myWriter
	mW.w = w
	mW.Write(b)

	return mW, &lenB
}

func (w myWriter) Write(b []byte) (int, error) {
	for _, v := range b {
		w.data = append(w.data, v)
	}
	return len(b), nil
}
