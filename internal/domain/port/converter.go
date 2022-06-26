package port

type Converter interface {
	Convert() (Messageable, error)
}
