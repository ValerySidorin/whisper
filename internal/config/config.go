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
	Provider  string
	Route     string
	Action    string
	Exporters []Exporter
}

type Exporter struct {
	Provider   string
	Attributes map[string]interface{}
}
