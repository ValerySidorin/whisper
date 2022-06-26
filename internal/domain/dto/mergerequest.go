package dto

import (
	"fmt"
	"time"
)

type MergeRequest struct {
	ID           int64
	IID          int64
	ProjectID    int64
	Title        string
	Description  string
	State        string
	CreatedDate  time.Time
	UpdatedDate  time.Time
	URL          string
	Author       Person
	Assignee     Person
	SourceBranch string
	TargetBranch string
	Labels       []string
}

func (mr *MergeRequest) GetMessage() string {
	return fmt.Sprintf("Your typical merge request message. Title: %v", mr.Title)
}
