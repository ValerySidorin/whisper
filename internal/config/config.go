package config

import (
	"time"

	"github.com/ValerySidorin/whisper/internal/domain/dto"
)

type Configuration struct {
	Http       HTTP
	VCSHosting VCSHosting
	Messenger  Messenger
	Storage    Storage
	Handlers   []Handler
}

type HTTP struct {
	Port    string
	Timeout int
}

func (h *HTTP) GetTimeout() time.Duration {
	return time.Duration(h.Timeout) * time.Second
}

type Handler struct {
	Route  string
	Action string
}

type Messenger struct {
	Provider dto.MessengerType
	Options  map[string]interface{}
}

type VCSHosting struct {
	Provider dto.VCSHostingType
	Options  map[string]interface{}
}

type Storage struct {
	Provider dto.StorageType
	Options  map[string]interface{}
}
