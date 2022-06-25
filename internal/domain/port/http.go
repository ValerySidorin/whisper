package port

type HttpError interface {
	error
	GetCode() int
	GetErr() error
}
