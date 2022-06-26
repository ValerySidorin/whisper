package port

type Exporter interface {
	SendMessage(msg string) error
}
