package dto

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

type MergeRequest struct {
	ID           int64
	IID          int64
	Project      Project
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
	trackerLink := ""
	matched, _ := regexp.MatchString(`^\w+-\d+\s`, mr.Title)
	if matched {
		trackerLink = "\nТрекер: https://tracker.yandex.ru/" + mr.Title[:strings.IndexByte(mr.Title, ' ')] + "\n"
	}
	switch mr.State {
	case "opened":
		return fmt.Sprintf("Новый Merge Request! | %v\n\n%v\n-- -- -- --\n%v\n\n%v\n%v\n%v → %v\nАвтор: %v",
			mr.Project.Name, mr.Title, mr.Description, mr.URL, trackerLink, mr.SourceBranch, mr.TargetBranch, mr.Author.Name)
	default:
		return fmt.Sprintf("Обновление статуса Merge Request! | %v\n-- -- -- --\n→ %v", mr.Project.Name, mr.State)
	}
}
