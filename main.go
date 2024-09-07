package main

import (
	"github.com/gin-gonic/gin"
	"offlinejudge/controller/contest"
	"offlinejudge/controller/participant"
	"offlinejudge/controller/status"
	"offlinejudge/controller/submission"
	"offlinejudge/io"
	"offlinejudge/utils"
)

func main() {

	r := gin.Default()

	participant.Init(r)
	submission.Init(r)
	status.Init(r)
	contest.Init(r)
	utils.SaveData = io.Update
	io.Init()

	r.OPTIONS("/*path", func(c *gin.Context) {
		o := c.Request.Header.Get("Origin")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", o)
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Header("Content-Type", "application/json")
		c.String(200, "")
	})

	err := r.Run(":12138")
	if err != nil {
		panic(err)
		return
	}
}
