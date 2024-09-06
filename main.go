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
	utils.SaveData = io.Update

	r := gin.Default()

	participant.Init(r)
	submission.Init(r)
	status.Init(r)
	contest.Init(r)
	io.Init()

	err := r.Run(":12138")
	if err != nil {
		panic(err)
		return
	}
}
