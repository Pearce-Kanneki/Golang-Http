package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/* Any */
func any(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"code":   100,
		"status": "ok",
	})
}
