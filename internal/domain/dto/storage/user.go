package storage

type User struct {
	ID                 int64  `json:"id"`
	VCSHostingID       int64  `json:"vcshosting_id"`
	MessengerID        int64  `json:"messenger_id"`
	VCSHostingUsername string `json:"vcshosting_username"`
	MessengerUsername  string `json:"messenger_username"`
}
