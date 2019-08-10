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
