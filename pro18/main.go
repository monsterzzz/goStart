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
		strLi := strings.SplitN(strData, " ", 2)
		fmt.Println(strLi)

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
			io.WriteString(c, user.cur+"\n")
		case "cd":
			if len(strLi) != 2 {
				io.WriteString(c, "error command!\n")
				continue
			}
			//fmt.Println(strLi)
			//fmt.Println(user.cur)
			f, err := os.Stat(user.cur + strLi[1])
			if err != nil {
				io.WriteString(c, "file not exits!")
				continue
			}
			if !f.IsDir() {
				io.WriteString(c, "dir not exits!\n")
				continue
			}
			if strLi[1] == ".." {
				if user.cur == "./" {
					continue
				}
				//fmt.Println(user.cur[:strings.LastIndexAny(user.cur[:len(user.cur) - 1],"/")])
				user.cur = user.cur[:strings.LastIndexAny(user.cur[:len(user.cur)-1], "/")]
				if user.cur == "." {
					user.cur += "/"
				}
			} else {
				user.cur += strLi[1] + "/"
			}
		case "cat":
			if len(strLi) != 2 {
				io.WriteString(c, "error command!\n")
				continue
			}
			//fmt.Println(strLi)
			//fmt.Println(user.cur)
			f, err := os.Stat(user.cur + strLi[1])
			if err != nil {
				io.WriteString(c, "file not exits!\n")
				continue
			}
			if f.IsDir() {
				io.WriteString(c, "the file name is a dir\n")
				continue
			}
			io.WriteString(c, user.cur+f.Name())
			fC, err := os.Open(user.cur + f.Name())
			defer fC.Close()
			if err != nil {
				io.WriteString(c, fmt.Sprint(err)+"\n")
			}
			b, err := ioutil.ReadAll(fC)
			if err != nil {
				io.WriteString(c, fmt.Sprint(err)+"\n")
			}
			io.WriteString(c, string(b))
		case "cp":
			if len(strLi) != 2 {
				io.WriteString(c, "error command!\n")
				continue
			}
			//fmt.Println(strLi)
			//fmt.Println(user.cur)
			f, err := os.Stat(user.cur + strLi[1])
			if err != nil {
				io.WriteString(c, "file not exits!\n")
				continue
			}
			if f.IsDir() {
				io.WriteString(c, " the file name is a dir\n")
				continue
			}
			block := make([]byte, 255)
			fC, err := os.Open(user.cur + strLi[1])
			defer fC.Close()
			for {
				if len(block) == 255 {
					block = make([]byte, 255)
				}
				n, err := fC.Read(block)
				if err == io.EOF {
					break
				}
				if n == 0 {
					break
				}
				c.Write(block)
			}
		}

	}
}
