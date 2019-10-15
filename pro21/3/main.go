package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	// {
	//    "message": "hello"
	//}

	r.Use(gin.Logger()) // 日志

	r.Use(gin.Recovery())

	r.GET("/hello", func(context *gin.Context) {
		//d := context.Query("ss")
		//fmt.Println(d[0])
		//panic("ss")
		context.JSON(200, gin.H{
			"message": "hello",
		})
	})
	r.Run()
}
