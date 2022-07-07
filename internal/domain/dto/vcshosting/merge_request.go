package vcshosting

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

type MergeRequestEvent struct {
	MergeRequest *MergeRequest
	Event        string
}
