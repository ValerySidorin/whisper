package port

type Converter interface {
	Convert() (interface{}, error)
}
