package dto

type Deployment struct {
	ObjectKind    string  `json:"object_kind"`
	Status        string  `json:"status"`
	DeploymentID  int64   `json:"deployment_id"`
	DeployableID  int64   `json:"deployable_id"`
	DeployableURL string  `json:"deployable_url"`
	Environment   string  `json:"environment"`
	Project       Project `json:"project"`
	ShortSha      string  `json:"short_sha"`
	User          User    `json:"user"`
	UserURL       string  `json:"user_url"`
	CommitURL     string  `json:"commit_url"`
	CommitTitle   string  `json:"commit_title"`
}
