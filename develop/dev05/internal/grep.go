package grep

import (
	"sync"
)

type Grep struct {
	config *Config
	wg     *sync.WaitGroup
}

func NewGrep() *Grep {
	return &Grep{
		wg:     &sync.WaitGroup{},
		config: &Config{},
	}
}

func (g *Grep) Init() error {
	if err := g.config.SetConfig(); err != nil {
		return err
	}
	return nil
}
func (g *Grep) Run() {
	for _, file := range g.config.Files {
		filter := NewFilter(file, g.config, g.wg)
		g.wg.Add(1)
		go filter.RunFilter()
	}
	if len(g.config.Files) < 1 {
		filter := NewFilter("", g.config, g.wg)
		g.wg.Add(1)
		filter.RunFilter()
	}
	g.wg.Wait()
}
