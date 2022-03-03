package cut

import (
	"errors"
	"flag"
	"strconv"
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
	c.FValues = make(FValues)
	return c
}

func (c *Config) ParseFlagF() error {
	sections := strings.Split(c.F, ",")
	for _, section := range sections {
		if strings.Contains(section, "-") {
			if len(section) > 1 && strings.Count(section, "-") == 1 {
				if section[0] == '-' {
					val, err := strconv.Atoi(section[1:])
					if err != nil {
						return err
					}
					for i := 0; i < val; i++ {
						c.FValues[i] = false
					}
				} else if section[len(section)-1] == '-' {
					val, err := strconv.Atoi(section[:len(section)-1])
					if err != nil {
						return err
					}
					val--
					c.FValues[val] = true
				} else {
					index := strings.Index(section, "-")
					valStart, err := strconv.Atoi(section[:index])
					if err != nil {
						return err
					}
					valEnd, err := strconv.Atoi(section[index+1:])
					if err != nil {
						return err
					}
					for i := valStart - 1; i < valEnd; i++ {
						c.FValues[i] = false
					}
				}
			} else {
				return errors.New("Incorrect Flag -f")
			}

		} else {
			val, err := strconv.Atoi(section)
			if err != nil {
				return err
			}
			val--
			c.FValues[val] = false
		}
	}
	return nil
}
