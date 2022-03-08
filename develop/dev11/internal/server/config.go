package server

type Config struct {
	host string
	port string
}

func NewConfig() *Config {
	return new(Config)
}

func (c *Config) Init(host string, port string) {
	c.host, c.port = host, port
}

func (c *Config) GetAddress() string {
	return c.host + ":" + c.port
}
