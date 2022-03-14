package grep

import (
	"sync"
)

// Grep Структура утилиты
type Grep struct {
	config *Config
	wg     *sync.WaitGroup
}

// NewGrep Создать Grep
func NewGrep() *Grep {
	return &Grep{
		wg:     &sync.WaitGroup{},
		config: &Config{},
	}
}

// Init Инициализация конфигурации утилиты аргументами
func (g *Grep) Init() error {
	if err := g.config.SetConfig(); err != nil {
		return err
	}
	return nil
}

// Run Запуск утилиты
func (g *Grep) Run() {
	for _, file := range g.config.Files {
		filter := NewFilter(file, g.config, g.wg) // Создаем фильтр на каждый файл
		g.wg.Add(1)
		go filter.RunFilter() // фильтрация в отдельной горутине
	}
	if len(g.config.Files) < 1 { // если файлы фильтрации не заданы, фильтруем по stdio
		filter := NewFilter("", g.config, g.wg)
		g.wg.Add(1)
		filter.RunFilter()
	}
	g.wg.Wait()
}
