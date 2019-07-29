package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	conn.Write([]byte("q"))

	//ipt :=bufio.NewScanner(os.Stdin)
	//for ipt.Scan(){
	//	n,err := conn.Write(ipt.Bytes())
	//	if err != nil{
	//		log.Fatal(err.Error())
	//	}
	//	fmt.Println(n)
	//}
	conn.Close()
}

func handleConn(c net.Conn) {
	defer func() {
		c.Close()
		fmt.Printf("client %s quit\n", c.RemoteAddr().String())
	}()
	for {
		data := make([]byte, 10)
		n, err := c.Read(data)
		if err == io.EOF {
			fmt.Printf("recive: %s \n", string(data[:n]))
			return
		}
	}
}
