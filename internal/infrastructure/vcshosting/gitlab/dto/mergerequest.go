package dto

import "time"

type MergeRequest struct {
	ObjectKind       string                 `json:"object_kind"`
	EventType        string                 `json:"event_type"`
	User             User                   `json:"user"`
	Project          Project                `json:"project"`
	Repository       Repository             `json:"repository"`
	ObjectAttributes MergeRequestAttributes `json:"object_attributes"`
}

type MergeRequestAttributes struct {
	ID                          int64     `json:"id"`
	TargetBranch                string    `json:"target_branch"`
	SourceBranch                string    `json:"source_branch"`
	SourceProjectID             int64     `json:"source_project_id"`
	AuthorID                    int64     `json:"author_id"`
	AssigneeID                  int64     `json:"assignee_id"`
	Title                       string    `json:"title"`
	CreatedAt                   time.Time `json:"created_at"`
	UpdatedAt                   time.Time `json:"updated_at"`
	MilestoneID                 int64     `json:"milestone_id"`
	State                       string    `json:"state"`
	BlockingDiscussionsResolved bool      `json:"blocking_discussions_resolved"`
	MergeStatus                 string    `json:"merge_status"`
	TargetProjectID             int64     `json:"target_project_id"`
	IID                         int64     `json:"iid"`
	Description                 string    `json:"description"`
	Source                      Project   `json:"source"`
	Target                      Project   `json:"target"`
	LastCommit                  Commit    `json:"last_commit"`
	WorkInProgress              bool      `json:"work_in_progress"`
	URL                         string    `json:"url"`
	Action                      string    `json:"action"`
	Assignee                    User      `json:"assignee"`
	Changes                     Changes   `json:"changes"`
}
