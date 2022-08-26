package vcshosting

type Build struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Stage       string  `json:"stage"`
	Status      string  `json:"status"`
	Duration    float64 `json:"duration"`
	Project     Project `json:"project"`
	User        Person  `json:"user"`
	Environment string  `json:"environment"`
}

type BuildEvent struct {
	Build *Build
}
