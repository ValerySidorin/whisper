package gitlabdto

type Project struct {
	Id                int64
	Name              string
	Description       string
	WebURL            string
	AvatarURL         string
	GitSSHURL         string
	GitHTTPURL        string
	Namespace         string
	VisibilityLevel   int64
	PathWithNamespace string
	DefaultBranch     string
	CIConfigPath      string
	Homepage          string
	URL               string
	SSHURL            string
	HTTPURL           string
}
