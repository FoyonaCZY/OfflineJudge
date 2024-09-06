package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var SaveData func() = nil

func Response(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "R U OK?",
		"data": data,
	})
}
