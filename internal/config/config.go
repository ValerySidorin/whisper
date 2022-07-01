package config

import "time"

type Configuration struct {
	Http       HTTP
	VCSHosting VCSHosting
	Messenger  Messenger
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
	Provider string
	Options  map[string]interface{}
}

type VCSHosting struct {
	Provider string
	Options  map[string]interface{}
}
