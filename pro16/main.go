package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

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
		handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer func() {
		fmt.Println("closing.. " + c.RemoteAddr().String())
		c.Close()
	}()
	for {
		_, err := io.WriteString(c, "hello,client\n")
		if err != nil {
			log.Print(err)
			return
		}
		time.Sleep(1 * time.Second)
	}
}
