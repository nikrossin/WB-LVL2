package utelnet

import (
	"flag"
	"log"
	"time"
)

// Config Конфигурация telent
type Config struct {
	Host    string
	Port    string
	Timeout time.Duration
}

// NewConfig Создать конфигурацию
func NewConfig() *Config {
	return &Config{}
}

// Init Инициализация конфигурации аргументами
func (c *Config) Init() {
	flag.DurationVar(&c.Timeout, "timeout", time.Second*10, "Значение timeout")
	flag.Parse()
	if len(flag.Args()) != 2 {
		log.Fatalln("count of parameters more or less than two")
	}
	c.Host, c.Port = flag.Arg(0), flag.Arg(1)
}
