package utelnet

import (
	"flag"
	"log"
	"time"
)

type Config struct {
	Host    string
	Port    string
	Timeout time.Duration
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Init() {
	flag.DurationVar(&c.Timeout, "timeout", time.Second*10, "Значение timeout")
	flag.Parse()
	if len(flag.Args()) != 2 {
		log.Fatalln("count of parameters more or less than two")
	}
	c.Host, c.Port = flag.Arg(0), flag.Arg(1)
}
