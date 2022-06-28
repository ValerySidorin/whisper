package dto

import "fmt"

type Deployment struct {
	ID            int64   `json:"id"`
	Status        string  `json:"status"`
	Job           Job     `json:"job"`
	DeployableURL string  `json:"deployable_url"`
	Environment   string  `json:"environment"`
	Project       Project `json:"project"`
	User          Person  `json:"user"`
	CommitURL     string  `json:"commit_url"`
	CommitTitle   string  `json:"commit_title"`
}

func (d *Deployment) GetMessage() string {
	switch d.Status {
	case "running":
		return fmt.Sprintf("A new deployment is running! | %v\n-- -- -- --\nEnv: %v\n-- -- -- --\nJob: %v\n %v\n-- -- -- --\nCommit: %v\n%v\n-- -- -- --\nInitiator: %v", d.Project.Name, d.Environment, d.Job.Name, d.DeployableURL, d.CommitTitle, d.CommitURL, d.User.Name)
	case "success":
		return fmt.Sprintf("Deployment succeded! | %v\n-- -- -- --\nEnv: %v\n-- -- -- --\nJob: %v\n-- -- -- --\nCommit: %v\n%v\n-- -- -- --\nInitiator: %v", d.Project.Name, d.Environment, d.Job.Name, d.CommitTitle, d.CommitURL, d.User.Name)
	default:
		return fmt.Sprintf("Deployment status changed! | %v\n-- -- -- --\nEnv: %v\n-- -- -- --\n→ %v\n-- -- -- --\nJob: %v\n%v", d.Project.Name, d.Environment, d.Status, d.Job.Name, d.DeployableURL)
	}
}
