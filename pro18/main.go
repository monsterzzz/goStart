package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
)

type User struct {
	c   net.Conn
	cur string
}

func (u *User) String() string {
	return fmt.Sprintf("{addr : %s , cur_pos : %s }", u.c.RemoteAddr().String(), u.cur)
}

func main() {
	listen, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal("1")
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal("2")
		}
		fmt.Println(conn.RemoteAddr().String())
		user := &User{c: conn, cur: "./"}
		handleConn(user)
	}
}

func handleConn(user *User) {
	c := user.c
	defer func() {
		c.Close()
		fmt.Printf("client %s quit\n", c.RemoteAddr().String())
	}()
	for {
		data := make([]byte, 255)
		n, err := c.Read(data)
		if err != nil {
			return
		}
		if n == 0 {
			continue
		}
		strData := strings.TrimSpace(string(data[:n]))
		fmt.Println(fmt.Sprintf("recive :  %s", strData))
		if strData == "q" {
			fmt.Println(">>>quit")
			return
		}
		strLi := strings.SplitN(strData, " ", 1)
		switch strLi[0] {
		case "ls":
			f, err := ioutil.ReadDir(user.cur)
			if err != nil {
				return
			}
			fName := make([]string, len(f))
			for i := 0; i < len(f); i++ {
				fName[i] = f[i].Name()
			}
			endStr := strings.Join(fName, "\n") + "\n"
			io.WriteString(c, endStr)
		case "cur":
			io.WriteString(c, user.cur)
		case "cd":
			if len(strLi) != 2 {
				io.WriteString(c, "error command!\n")
				continue
			}
			fmt.Println(user.cur)
			f, err := os.Stat(user.cur + "/" + strLi[1])
			if err != nil {
				io.WriteString(c, fmt.Sprint(err))
				continue
			}
			if !f.IsDir() {
				io.WriteString(c, "dir not exits!\n")
				continue
			}
			user.cur = strLi[1]
		}

	}
}
