package main

import (
	"fmt"
	"io"
)

type S struct {
	mov  int
	data []byte
}

func main() {
	//f,err := os.Open("a.txt")
	//if err!= nil{
	//	log.Fatal("sss")
	//}
	//
	//ioutil.ReadAll()

	//var s S
	//s = S{data:[]byte("ssss")}
	//data,err := ioutil.ReadAll(&s)
	//if err != nil{
	//	log.Fatal("s2")
	//}
	////for i,v := range s.data{
	////	fmt.Printf("%d -> %s",i,v)
	////}
	//fmt.Println(data)

	b := make([]byte, 2)
	s := S{data: []byte("1234")}

	for i := 0; i < len(s.data); i++ {
		_, err := s.Read(b)
		if err == io.EOF {
			break
		}
		fmt.Println(string(b))
	}
}

func (s *S) Read(b []byte) (int, error) {
	var n = 0
	if len(b) >= len(s.data) {
		for i := 0; i < len(s.data); i++ {
			b[i] = s.data[i]
			n++
		}
		return n, io.EOF
	}
	for i := 0; i < len(b); i++ {
		//fmt.Println(s.mov + i)
		//
		cur := s.mov + i
		if s.mov != 0 {
			cur += 1
		}
		if cur == len(s.data) {
			return 0, io.EOF
		}

		b[i] = s.data[cur]
		n++
	}
	s.mov += len(b) - 1
	return n, nil

}
