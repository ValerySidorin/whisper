package port

type HttpError interface {
	error
	GetCode() int
	GetErr() error
}

type VCSHostingHandler interface {
	HandleMergeRequest(body []byte) error
	HandleDeployment(body []byte) error
}
