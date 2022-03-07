package wget

import (
	"flag"
	"log"
	"net/url"
)

type Config struct {
	Domain *url.URL
	Dir    string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Init() {
	var uri string
	flag.StringVar(&uri, "u", "", "URL-адрес сайта")
	flag.StringVar(&c.Dir, "p", ".", "Путь до каталога сохранения")
	flag.Parse()
	var err error
	c.Domain, err = url.ParseRequestURI(uri)
	if err != nil {
		log.Fatalln(err)
	}
}
