package dto

import "fmt"

type Deployment struct {
	ID            int64   `json:"id"`
	Status        string  `json:"status"`
	DeployableURL string  `json:"deployable_url"`
	Environment   string  `json:"environment"`
	Project       Project `json:"project"`
	User          Person  `json:"user"`
	CommitURL     string  `json:"commit_url"`
	CommitTitle   string  `json:"commit_title"`
}

func (d *Deployment) GetMessage() string {
	switch d.Status {
	case "success":
		return fmt.Sprintf("Произошел успешный Deploy! | %v\n\nОкружение: %v\n\nЗадача: %v\n\nКоммит: %v\n%v\n\nИнициатор: %v", d.Project.Name, d.Environment, d.DeployableURL, d.CommitTitle, d.CommitURL, d.User.Name)
	default:
		return fmt.Sprintf("Изменение статуса Deploy! | %v\n\nОкружение: %v\n\n→ %v\n\nЗадача: %v", d.Project.Name, d.Environment, d.Status, d.DeployableURL)
	}
}
