package main

import (
	"fmt"
	"net/http"

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
	router.POST("/post", postFun)

	router.Run(":8080")
}

func getFun(context *gin.Context) {
	result.Code = 100
	result.Message = "Success Get"

	context.JSON(http.StatusOK, result)
}

func getFun2(context *gin.Context) {
	input := context.Query("q")
	result.Code = 101
	result.Message = "Message:" + input

	context.JSON(http.StatusOK, result)
}

func postFun(context *gin.Context) {
	result.Code = 100
	result.Message = "Success POST"

	context.JSON(http.StatusOK, result)
}
