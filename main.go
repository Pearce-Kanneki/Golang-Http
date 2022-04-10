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
	router.GET("/p1", getFun)
	router.GET("/p2", getFun2)
	router.GET("/p3/:input", getFun3)

	router.POST("/post", postFun)
	router.POST("/post/v1", postFun2)

	router.Run(":8080")
}
