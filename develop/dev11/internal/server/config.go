package server

// Config Конфигурция сервера
type Config struct {
	host string
	port string
}

// NewConfig Создание конфигурации сервера
func NewConfig() *Config {
	return new(Config)
}

// Set Установить значения адреса и порта
func (c *Config) Set(host string, port string) {
	c.host, c.port = host, port
}

// DefaultSet Установить значения для адреса и порта по умолчанию
func (c *Config) DefaultSet() {
	c.host, c.port = "localhost", "8080"
}

// GetAddress Получить полный адрес
func (c *Config) GetAddress() string {
	return c.host + ":" + c.port
}
