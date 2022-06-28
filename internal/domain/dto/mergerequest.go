package dto

import (
	"fmt"
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
	switch mr.State {
	case "opened":
		return fmt.Sprintf("New Merge Request! | %v\n-- -- -- --\nTitle: %v\nDescription: %v\n%v\n-- -- -- --\nBranch: %v → %v\nAuthor: %v",
			mr.Project.Name, mr.Title, mr.Description, mr.URL, mr.SourceBranch, mr.TargetBranch, mr.Author.Name)
	case "merged":
		return fmt.Sprintf("Merge Request has been merged! | %v\n-- -- -- --\nTitle: %v\nDescription: %v\n%v\n-- -- -- --\nBranch: %v → %v\nAuthor: %v",
			mr.Project.Name, mr.Title, mr.Description, mr.URL, mr.SourceBranch, mr.TargetBranch, mr.Author.Name)
	default:
		return fmt.Sprintf("Merge Request status changed! | %v\n-- -- -- --\n→ %v", mr.Project.Name, mr.State)
	}
}
