package contest

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"offlinejudge/utils"
)

type Info struct { //比赛信息
	Count int `json:"count"`
	Begin int `json:"begin"`
	End   int `json:"end"`
}

var (
	Contest Info
)

func SetContestInfo(count, begin, end int) {
	Contest = Info{
		Count: count,
		Begin: begin,
		End:   end,
	}
	utils.SaveData()
}

func GetContestInfo() Info {
	return Contest
}

func Init(r *gin.Engine) {
	r.POST("/api/contest/setContestInfo", func(c *gin.Context) {
		var info Info
		err := c.BindJSON(&info)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		SetContestInfo(info.Count, info.Begin, info.End)
		utils.Response(c, nil)
	})

	r.GET("/api/contest/getContestInfo", func(c *gin.Context) {
		utils.Response(c, GetContestInfo())
	})
}
