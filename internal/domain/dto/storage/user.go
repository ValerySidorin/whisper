package storage

type User struct {
	ID                 int64  `json:"id"`
	VCSHosting         string `json:"vcshosting"`
	Messenger          int64  `json:"messenger"`
	VCSHostingUsername string `json:"vcshosting_username"`
	MessengerUsername  string `json:"messenger_username"`
	State              int64  `json:"state"`
}
