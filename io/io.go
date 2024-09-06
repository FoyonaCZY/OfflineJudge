package io

import (
	"encoding/json"
	"offlinejudge/controller/contest"
	"offlinejudge/controller/participant"
	"offlinejudge/controller/submission"
	"os"
)

func input() error {
	// 输入
	res := struct {
		SubmissionsID []int                           `json:"submissionsId"`
		Submissions   map[int]submission.Submission   `json:"submissions"`
		SubId         int                             `json:"submissionId"`
		PartId        int                             `json:"participantId"`
		Users         map[int]participant.Participant `json:"users"`
		UsersID       []int                           `json:"usersId"`
		Contest       contest.Info                    `json:"contest"`
	}{}
	a, err := os.ReadFile("./data/data.json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(a, &res)
	submission.SubmissionsID = res.SubmissionsID
	submission.Submissions = res.Submissions
	submission.SubId = res.SubId
	participant.PartId = res.PartId
	participant.Users = res.Users
	participant.UsersID = res.UsersID
	contest.Contest = res.Contest
	return err
}

func output() error {
	// 输出
	file, err := os.Create("./data/data.json")
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	res := struct {
		SubmissionsID []int                           `json:"submissionsId"`
		Submissions   map[int]submission.Submission   `json:"submissions"`
		SubId         int                             `json:"submissionId"`
		PartId        int                             `json:"participantId"`
		Users         map[int]participant.Participant `json:"users"`
		UsersID       []int                           `json:"usersId"`
		Contest       contest.Info                    `json:"contest"`
	}{
		SubmissionsID: submission.SubmissionsID,
		Submissions:   submission.Submissions,
		SubId:         submission.SubId,
		PartId:        participant.PartId,
		Users:         participant.Users,
		UsersID:       participant.UsersID,
		Contest:       contest.Contest,
	}
	a, _ := json.Marshal(res)
	_, _ = file.Write(a)
	return nil
}

func Init() {
	// 初始化
	_ = input()
}

func Update() {
	// 更新
	_ = output()
}
