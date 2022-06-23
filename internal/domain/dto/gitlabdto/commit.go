package gitlabdto

import "time"

type Commit struct {
	Id        []byte
	Message   string
	Title     string
	Timestamp time.Time
	URL       string
	Author    Author
	Added     []string
	Modified  []string
	Removed   []string
}
