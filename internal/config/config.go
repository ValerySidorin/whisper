package config

type Configuration struct {
	HTTP   HTTP
	Routes Routes
}

type HTTP struct {
	Port    string
	Timeout int
}

type Routes struct {
	GitlabRoutes []string
}
