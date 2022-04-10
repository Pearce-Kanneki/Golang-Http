package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/* GET */
func getFun(context *gin.Context) {
	result.Code = 100
	result.Message = "Success Get"

	context.JSON(http.StatusOK, result)
}

/* GET 查詢參數(Query Params) */
func getFun2(context *gin.Context) {
	input := context.Query("q")
	result.Code = 101
	result.Message = "Message:" + input

	context.JSON(http.StatusOK, result)
}

/* GET 路徑參數(Path Params) */
func getFun3(context *gin.Context) {
	msg := context.Param("input")
	result.Code = 101
	result.Message = "Message type:" + msg

	context.JSON(http.StatusOK, result)
}
