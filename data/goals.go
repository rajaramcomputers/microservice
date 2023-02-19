package data

import (
	"encoding/json"
	"io"
)

type Goal struct {
	GoalID           int    `json:"id"`
	GoalDetail       string `json:"goaldetail"`
	TaskID           int    `json:"taskid"`
	TaskName         string `json:"taskname"`
	BelowExpectation string `json:"belowexpectation"`
	MeetExpectation  string `json:"meetexpectation"`
	AboveExpectation string `json:"aboveexpectation"`
}

type Goals []*Goal

func (g *Goals) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(g)
}

func GetGoals() Goals {
	return goalList
}

var goalList = []*Goal{
	{
		GoalID:           1,
		GoalDetail:       "to improve work plan",
		TaskID:           1,
		TaskName:         "to increase the research",
		BelowExpectation: "10-04-2023",
		MeetExpectation:  "01-04-2023",
		AboveExpectation: "05-04-2023",
	},
	{
		GoalID:           1,
		GoalDetail:       "to improve work plan",
		TaskID:           2,
		TaskName:         "to increase the colloboration",
		BelowExpectation: "10-04-2023",
		MeetExpectation:  "01-04-2023",
		AboveExpectation: "05-04-2023",
	},
}
