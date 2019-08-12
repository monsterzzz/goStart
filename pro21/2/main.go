package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gpmgo/gopm/modules/log"
	"runtime"
)

func main() {
	r := gin.Default()

	r.GET("/user0/:name/:sex", pathParam)
	r.GET("/user1/", getQueryParam)
	r.POST("/user2", postQueryParam)
	r.POST("/user3", postGetParamBody)
	r.POST("/user4", postFile)
	r.POST("/user5", manyPostFile)
	r.Run()
}

func pathParam(c *gin.Context) {
	funcName, _, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("!")
	}
	name := c.Param("name")
	sex := c.Param("sex")
	message := fmt.Sprintf("hello,%s,%s,welcome to our %s Test!", name, sex, runtime.FuncForPC(funcName).Name())
	c.String(200, message)
}

func getQueryParam(c *gin.Context) {
	funcName, _, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("!")
	}
	name := c.DefaultQuery("name", "guest")
	sex := c.DefaultQuery("sex", "unknow")
	message := fmt.Sprintf("hello,%s,%s,welcome to our %s Test!", name, sex, runtime.FuncForPC(funcName).Name())
	c.String(200, message)
}

func postQueryParam(c *gin.Context) {
	funcName, _, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("!")
	}

	name := c.PostForm("name")
	//s := c.PostForm("name")
	//fmt.Println(s)
	sex := c.DefaultPostForm("sex", "unknow")
	message := fmt.Sprintf("hello,%s,%s,welcome to our %s Test!", name, sex, runtime.FuncForPC(funcName).Name())
	c.String(200, message)
}

func postGetParamBody(c *gin.Context) {
	funcName, _, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("!")
	}

	name := c.DefaultQuery("name", "ww1")
	//s := c.PostForm("name")
	//fmt.Println(s)
	sex := c.DefaultPostForm("sex", "unknow")
	message := fmt.Sprintf("hello,%s,%s,welcome to our %s Test!", name, sex, runtime.FuncForPC(funcName).Name())
	c.String(200, message)
}

func postFile(c *gin.Context) {
	funcName, _, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("!")
	}

	//name := c.DefaultQuery("name","ww1")
	////s := c.PostForm("name")
	////fmt.Println(s)
	//sex := c.DefaultPostForm("sex", "unknow")
	file, _ := c.FormFile("file")
	fmt.Println(file.Filename)
	e := c.SaveUploadedFile(file, "pro21/2/hh.txt")
	fmt.Println(e)
	message := fmt.Sprintf("hello,welcome to our %s Test!,upload success!", runtime.FuncForPC(funcName).Name())
	c.String(200, message)
	//curl -X POST http://localhost:8080/user4 -F "file=@fileD/hello.txt" -H "Content-Type: multipart/form-data"
}

func manyPostFile(c *gin.Context) {
	funcName, _, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("!")
	}

	//name := c.DefaultQuery("name","ww1")
	////s := c.PostForm("name")
	////fmt.Println(s)
	//sex := c.DefaultPostForm("sex", "unknow")
	form, e := c.MultipartForm()
	fmt.Println(e)
	files := form.File["upload[]"]
	fmt.Println(form.File)
	fmt.Println(len(files))
	n := 0
	for _, file := range files {
		e := c.SaveUploadedFile(file, fmt.Sprintf("pro21/2/hh_many1_%d.txt", n))
		//e:= c.SaveUploadedFile(file,"pro21/2/hh_many_1d.txt")
		fmt.Println("sss")
		fmt.Println(e)
		n++
	}
	message := fmt.Sprintf("hello,welcome to our %s Test!,upload success!", runtime.FuncForPC(funcName).Name())
	c.String(200, message)
	//curl -X POST http://localhost:8080/user4 -F "file=@fileD/hello.txt" -H "Content-Type: multipart/form-data"
}
