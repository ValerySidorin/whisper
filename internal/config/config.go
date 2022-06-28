package config

import "time"

type Configuration struct {
	Http     HTTP
	Handlers []Handler
}

type HTTP struct {
	Port    string
	Timeout int
}

func (h *HTTP) GetTimeout() time.Duration {
	return time.Duration(h.Timeout) * time.Second
}

type Handler struct {
	VCSHosting VCSHosting
	Route      string
	Action     string
	Exporters  []Exporter
}

type Exporter struct {
	Provider string
	Options  map[string]interface{}
}

type VCSHosting struct {
	Provider string
	Options  map[string]interface{}
}
