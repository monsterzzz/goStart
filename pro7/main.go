package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type User struct {
	name  string
	repo  string
	token string
}

type Github interface {
	startGithub()
	getIssues()
	addIssues()
	changeIssues()
	closeIssues()
}

func (u *User) startGithub() {
	githubStart := "https://api.github.com/zen"
	resp := send_req(githubStart)
	fmt.Println(resp)
}

func (u *User) getIssues() {
	resp := send_req("https://api.github.com/repos/" + u.name + "/" + u.repo + "/issues?" + u.token)
	fmt.Println(resp)
}

func (u *User) addIssues(title, body string) {
	//add issues
	type My_json struct {
		Title string
		Body  string
	}
	var myj My_json
	myj.Title = title
	myj.Body = body
	j, err := json.Marshal(myj)
	if err != nil {
		os.Exit(1)
	}
	newJ := strings.ToLower(string(j))
	newB := []byte(newJ)
	resp := send_post("https://api.github.com/repos/"+u.name+"/"+u.repo+"/issues?"+u.token, newB)
	fmt.Println(resp)
}

func (u *User) changeIssues(issueNum int, title, body, state string) {
	type MyJson struct {
		Title string
		Body  string
		State string
	}
	var myj MyJson
	myj.Title = title
	myj.Body = body
	myj.State = state
	j, err := json.Marshal(myj)
	if err != nil {
		os.Exit(1)
	}
	newJ := strings.ToLower(string(j))
	newB := []byte(newJ)
	client := &http.Client{}
	req, err := http.NewRequest("PATCH", "https://api.github.com/repos/"+u.name+"/"+u.repo+"/issues/"+strconv.Itoa(issueNum)+"?"+u.token, bytes.NewReader(newB))
	if err != nil {
		os.Exit(1)
	}
	resp, err := client.Do(req)
	if err != nil {
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(b))

}

func main() {
	// github 开发
	var user User
	user.repo = "awesomeProject2"
	user.name = "monsterzzz"
	user.token = "access_token=ed57053e65050d8830d60a0500e734296434f5ef"

	//user.startGithub()
	//user.getIssues()
	//user.changeIssues(1,"hh","xxs","open")
}

func send_req(s string) string {
	resp, err := http.Get(s)
	if err != nil {
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Println("error input stream")
	}
	return fmt.Sprintf("%s", b)
}

func send_post(url string, data []byte) string {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		os.Exit(1)
	}
	resp, err := client.Do(req)
	if err != nil {
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(resp.Body)
	return fmt.Sprintf("%s", b)

}
