package server

type Config struct {
	host string
	port string
}

func NewConfig() *Config {
	return new(Config)
}

func (c *Config) Set(host string, port string) {
	c.host, c.port = host, port
}

func (c *Config) DefaultSet() {
	c.host, c.port = "localhost", "8080"
}
func (c *Config) GetAddress() string {
	return c.host + ":" + c.port
}
