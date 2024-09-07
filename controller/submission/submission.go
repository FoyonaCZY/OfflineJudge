package submission

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"offlinejudge/utils"
)

type Submission struct { //提交
	ID            int    `json:"id"`
	ParticipantID int    `json:"participantId"`
	ProblemID     int    `json:"problemId"`
	Time          int    `json:"time"`
	Status        string `json:"status"`
}

var (
	SubId         int
	Submissions   map[int]Submission
	SubmissionsID []int
)

func createSubmission(s Submission) int {
	//创建提交
	SubId++
	s.ID = SubId
	SubmissionsID = append(SubmissionsID, s.ID)
	Submissions[s.ID] = s
	utils.SaveData()
	return s.ID
}

func batchRemoveSubmission(deleteId []int) int {
	//批量删除提交
	res := 0
	for _, id := range deleteId {
		_, ok := Submissions[id]
		if ok {
			res++
			delete(Submissions, id)
			for i, v := range SubmissionsID {
				if v == id {
					SubmissionsID = append(SubmissionsID[:i], SubmissionsID[i+1:]...)
					break
				}
			}
		}
	}
	utils.SaveData()
	return res
}

func listSubmissions() []Submission {
	//列出提交
	res := make([]Submission, 0)
	for _, v := range SubmissionsID {
		res = append(res, Submissions[v])
	}
	return res
}

func batchUpdateSubmission(a []Submission) int {
	//批量更新提交
	res := 0
	for _, s := range a {
		_, ok := Submissions[s.ID]
		if ok {
			res++
			tmp := Submissions[s.ID]
			tmp.Time = s.Time
			tmp.Status = s.Status
			Submissions[s.ID] = tmp
		}
	}
	utils.SaveData()
	return res
}

func Init(r *gin.Engine) {
	//初始化提交
	Submissions = make(map[int]Submission)

	r.POST("/api/submissions/create", func(c *gin.Context) {
		var s Submission
		err := c.BindJSON(&s)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res := struct {
			ID int `json:"id"`
		}{createSubmission(s)}
		utils.Response(c, res)
	})

	r.POST("/api/submissions/batchUpdate", func(c *gin.Context) {
		var a []Submission
		err := c.BindJSON(&a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res := struct {
			Count int `json:"count"`
		}{batchUpdateSubmission(a)}
		utils.Response(c, res)
	})

	r.POST("/api/submissions/batchRemove", func(c *gin.Context) {
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
		}{batchRemoveSubmission(res.Ids)}
		utils.Response(c, a)
	})

	r.GET("/api/submissions/list", func(c *gin.Context) {
		utils.Response(c, listSubmissions())
	})
}
