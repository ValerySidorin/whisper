package gitlabdto

type Push struct {
	ObjectKind        string
	EventName         string
	Before            []byte
	After             []byte
	Ref               string
	CheckoutSHA       []byte
	UserId            int64
	UserName          string
	UserUsername      string
	UserEmail         string
	UserAvatar        string
	ProjectId         int
	Project           Project
	Repository        Repository
	Commits           []Commit
	TotalCommitsCount int64
}
