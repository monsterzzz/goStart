package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

var n = 0

func main() {

	listen, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal("server start error")
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal("abort!")
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer func() {
		fmt.Println("closing.. " + c.RemoteAddr().String())
		c.Close()
	}()
	for {
		n++
		_, err := io.WriteString(c, fmt.Sprintf("hello,client!It is Server's %d message\n", n))
		if err != nil {
			log.Print(err)
			return
		}
		time.Sleep(1 * time.Second)
	}
}
