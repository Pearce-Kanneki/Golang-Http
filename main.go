package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var result = Result{}

func main() {
	fmt.Println("Golang Gin Http")

	router := gin.Default()
	router.RedirectFixedPath = true // URL 大小寫通吃
	router.GET("/p1", getFun)
	router.GET("/p2", getFun2)
	router.GET("/p3/:input", getFun3)

	router.POST("/post", postFun)
	router.POST("/post/v1", postFun2)

	router.Any("/any", any)

	/** 其他類型
	 * PUT, PATCH, HEAD, OPTIONS, DELETE, CONNECT, TRACE
	 */

	router.Run(":8080")
}
