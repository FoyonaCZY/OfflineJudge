package participant

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"offlinejudge/utils"
)

type Participant struct { //参赛者
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Department string `json:"department"`
}

var (
	PartId  int
	Users   map[int]Participant
	UsersID []int
)

func batchCreateParticipants(a []Participant) []int {
	var res []int
	for _, p := range a {
		PartId++
		p.ID = PartId
		UsersID = append(UsersID, PartId)
		Users[PartId] = p
		res = append(res, PartId)
	}
	utils.SaveData()
	return res
}

func batchUpdateParticipants(a []Participant) int {
	res := 0
	for _, p := range a {
		_, ok := Users[p.ID]
		if ok {
			res++
			Users[p.ID] = Participant{
				ID:         p.ID,
				Name:       p.Name,
				Department: p.Department,
			}
		}
	}
	utils.SaveData()
	return res
}

func batchRemoveParticipants(deleteId []int) int {
	res := 0
	for _, id := range deleteId {
		_, ok := Users[id]
		if ok {
			res++
			delete(Users, id)
			for i, v := range UsersID {
				if v == id {
					UsersID = append(UsersID[:i], UsersID[i+1:]...)
					break
				}
			}
		}
	}
	utils.SaveData()
	return res
}

func listAllParticipants() []Participant {
	res := make([]Participant, 0)
	for _, p := range UsersID {
		res = append(res, Users[p])
	}
	return res
}

func Init(r *gin.Engine) {
	Users = make(map[int]Participant)
	r.POST("/api/participants/batchCreate", func(c *gin.Context) { //批量创建参赛者
		var participants []Participant
		err := c.BindJSON(&participants)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		utils.Response(c, batchCreateParticipants(participants))
	})

	r.POST("/api/participants/batchRemove", func(c *gin.Context) { //批量删除参赛者
		res := struct {
			Ids []int `json:"ids"`
		}{}
		err := c.BindJSON(&res)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		a := struct {
			Count int `json:"count"`
		}{batchRemoveParticipants(res.Ids)}
		utils.Response(c, a)
	})

	r.POST("/api/participants/batchUpdate", func(c *gin.Context) { //批量更新参赛者
		var participants []Participant
		err := c.BindJSON(&participants)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		a := struct {
			Count int `json:"count"`
		}{batchUpdateParticipants(participants)}
		utils.Response(c, a)
	})

	r.GET("/api/participants/listAll", func(c *gin.Context) { //获取参赛者数据
		utils.Response(c, listAllParticipants())
	})
}
