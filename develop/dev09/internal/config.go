package wget

import (
	"flag"
	"log"
	"net/url"
)

// Config Структура с параметрами конфигурации
type Config struct {
	Domain *url.URL
	Dir    string
}

// NewConfig Создние конфига
func NewConfig() *Config {
	return &Config{}
}

// Init Инициализация флагов
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
