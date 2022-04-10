package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/* POST */
func postFun(context *gin.Context) {
	result.Code = 100
	result.Message = "Success POST"

	context.JSON(http.StatusOK, result)
}

/* POST Body */
func postFun2(c *gin.Context) {
	// msg := c.PostForm("q")
	msg := c.DefaultPostForm("q", "no value") // 沒輸入參數時,帶入預設值

	result.Code = 101
	result.Message = "POST Message:" + msg

	c.JSON(http.StatusOK, result)
}
