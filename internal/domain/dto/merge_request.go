package dto

import (
	"fmt"
	"regexp"
	"strings"
)

type MergeRequest struct {
	ID           int64
	IID          int64
	Project      Project
	Title        string
	Description  string
	State        string
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
		trackerLink = "https://tracker.yandex.ru/" + mr.Title[:strings.IndexByte(mr.Title, ' ')]
	}

	switch mr.State {
	case "opened":
		return fmt.Sprintf("New Merge Request! | %v\n-- -- -- --\nTitle: %v\nDescription: %v\n%v\n-- -- -- --\nTracker: %v\n-- -- -- --\nBranch: %v → %v\nAuthor: %v",
			mr.Project.Name, mr.Title, mr.Description, mr.URL, trackerLink, mr.SourceBranch, mr.TargetBranch, mr.Author.Name)
	default:
		return ""
	}
}
