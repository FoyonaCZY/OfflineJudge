package status

import (
	"github.com/gin-gonic/gin"
	"offlinejudge/controller/contest"
	"offlinejudge/controller/participant"
	"offlinejudge/controller/submission"
	"offlinejudge/utils"
	"slices"
	"strings"
)

type RecordSubmission struct { //记录提交
	ProblemID int    `json:"problemId"`
	Tries     int    `json:"tries"`
	Time      int    `json:"time"`
	Status    string `json:"status"`
}

type Record struct { //记录
	ParticipantID         int                `json:"participantId"`
	ParticipantName       string             `json:"participantName"`
	ParticipantDepartment string             `json:"participantDepartment"`
	Penalty               int                `json:"penalty"`
	Solved                int                `json:"solved"`
	Submissions           []RecordSubmission `json:"submissions"`
}
type BoardStatus struct { //榜单状态
	ProblemCount int      `json:"problemCount"`
	Begin        int      `json:"begin"`
	End          int      `json:"end"`
	Records      []Record `json:"records"`
}
type Pair struct {
	first  int
	second int
}

var (
	boardStatus BoardStatus
)

func updateBoardStatus() {
	//更新榜单数据
	slices.SortFunc(submission.SubmissionsID, func(i, j int) int {
		return submission.Submissions[i].Time - submission.Submissions[j].Time
	})
	recordSubmissions := make(map[Pair]RecordSubmission) //记录提交
	for _, i := range submission.SubmissionsID {
		s := submission.Submissions[i]
		pid := s.ProblemID
		uid := s.ParticipantID
		time := s.Time
		status := s.Status

		if recordSubmissions[Pair{uid, pid}].Status == "Accepted" { //已经AC
			continue
		}
		sub := RecordSubmission{
			ProblemID: pid,
			Tries:     recordSubmissions[Pair{uid, pid}].Tries + 1,
			Time:      time,
			Status:    "Accepted",
		}
		if status != "Accepted" {
			sub.Time = 0
			sub.Status = "Rejected"
		}
		recordSubmissions[Pair{uid, pid}] = sub
	}

	var records []Record
	for _, p := range participant.Users { //遍历参赛者
		penalty := 0
		solved := 0
		var submissions []RecordSubmission
		for pid := 1; pid <= contest.Contest.Count; pid++ {
			if recordSubmissions[Pair{p.ID, pid}].Status == "Accepted" {
				penalty += recordSubmissions[Pair{p.ID, pid}].Time + (recordSubmissions[Pair{p.ID, pid}].Tries-1)*20
				solved++
			}
			tmp := recordSubmissions[Pair{p.ID, pid}]
			tmp.ProblemID = pid
			recordSubmissions[Pair{p.ID, pid}] = tmp
			submissions = append(submissions, recordSubmissions[Pair{p.ID, pid}])
		}
		records = append(records, Record{
			ParticipantID:         p.ID,
			ParticipantName:       p.Name,
			ParticipantDepartment: p.Department,
			Penalty:               penalty,
			Solved:                solved,
			Submissions:           submissions,
		})
	}

	slices.SortFunc(records, func(i, j Record) int { //排序
		if i.Solved != j.Solved {
			return j.Solved - i.Solved
		}
		if i.Penalty != j.Penalty {
			return i.Penalty - j.Penalty
		}
		return strings.Compare(i.ParticipantName, j.ParticipantName)
	})

	info := contest.GetContestInfo()
	boardStatus = BoardStatus{
		ProblemCount: info.Count,
		Begin:        info.Begin,
		End:          info.End,
		Records:      records,
	}
}

func getBoardStatus() BoardStatus {
	//返回当前榜单
	updateBoardStatus()
	return boardStatus
}

func Init(r *gin.Engine) {
	r.GET("/api/board/status", func(c *gin.Context) { //获取榜单数据
		utils.Response(c, getBoardStatus())
	})
}
