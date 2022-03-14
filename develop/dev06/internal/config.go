package cut

import (
	"errors"
	"flag"
	"strconv"
	"strings"
)

// FValues Значения флага -f
type FValues map[int]bool

// Config Конфигурация cut
type Config struct {
	F       string
	D       string
	S       bool
	FValues FValues
}

// NewConfigInit Создание новой конфигурации с инициализацией
func NewConfigInit() *Config {
	c := &Config{}
	flag.StringVar(&c.F, "f", "", "выбрать поля (колонки)")
	flag.StringVar(&c.D, "d", "\t", "использовать другой разделитель")
	flag.BoolVar(&c.S, "s", false, "выводить только строки с разделителем")
	flag.Parse()
	c.FValues = make(FValues)
	return c
}

// ParseFlagF Анализ флага F для парсинга всех возможных значений
func (c *Config) ParseFlagF() error {
	sections := strings.Split(c.F, ",") // разбиваем на секции значений
	for _, section := range sections {
		if strings.Contains(section, "-") { // если есть интервал значений
			if len(section) > 1 && strings.Count(section, "-") == 1 { // проверка на один интревал в секции и хотя бы 1 значение
				if section[0] == '-' { // левый интервал
					val, err := strconv.Atoi(section[1:])
					if err != nil {
						return err
					}
					for i := 0; i < val; i++ {
						c.FValues[i] = false // добавляем в мапу все значения с 0 до значения после тире
					}
				} else if section[len(section)-1] == '-' { // если правый интервал
					val, err := strconv.Atoi(section[:len(section)-1])
					if err != nil {
						return err
					}
					val--
					c.FValues[val] = true // значение начала правого интервала true - интервал до конца строки
				} else { // интервал с заданными 2-мя значениями
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

		} else { // если секция без интервала
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
