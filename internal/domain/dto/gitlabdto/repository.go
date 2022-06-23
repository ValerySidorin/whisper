package gitlabdto

type Repository struct {
	Name            string
	URL             string
	Description     string
	Homepage        string
	GitHTTPURL      string
	GitSSHURL       string
	VisibilityLevel int64
}
