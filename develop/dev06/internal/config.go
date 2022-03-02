package cut

import (
	"flag"
	"strings"
)

type FValues map[int]bool

type Config struct {
	F       string
	D       string
	S       bool
	FValues FValues
}

func NewConfigInit() *Config {
	c := &Config{}
	flag.StringVar(&c.F, "f", "", "выбрать поля (колонки)")
	flag.StringVar(&c.D, "d", "\t", "использовать другой разделитель")
	flag.BoolVar(&c.S, "s", false, "выводить только строки с разделителем")
	flag.Parse()
	return c
}

func (c *Config) ParseFlagF() {
	sections := strings.Split(c.F, ",")
	for _, section := range sections {

	}
}
