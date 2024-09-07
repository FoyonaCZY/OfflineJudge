package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var SaveData func() = nil

func Response(c *gin.Context, data interface{}) {
	o := c.Request.Header.Get("Origin")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Origin", o)
	c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "R U OK?",
		"data": data,
	})
}
