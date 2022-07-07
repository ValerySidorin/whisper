package vcshosting

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

type DeploymentEvent struct {
	Deployment *Deployment
}
