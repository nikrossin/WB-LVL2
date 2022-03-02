package grep

import (
	"errors"
	"flag"
)

type Flags struct {
	AA int
	BB int
	CC int
	C  bool
	I  bool
	V  bool
	F  bool
	N  bool
}
type Config struct {
	Flags
	Pattern string
	Files   []string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) SetConfig() error {
	flag.IntVar(&c.AA, "A", 0, "печатать +N строк после совпадения")
	flag.IntVar(&c.BB, "B", 0, "печатать +N строк до совпадения")
	flag.IntVar(&c.CC, "C", 0, "печатать ±N строк вокруг совпадения")
	flag.BoolVar(&c.C, "c", false, "печатать количество строк")
	flag.BoolVar(&c.I, "i", false, "игнорировать регистр")
	flag.BoolVar(&c.V, "v", false, "вместо совпадения, исключать")
	flag.BoolVar(&c.F, "F", false, "точное совпадение со строкой, не паттерн")
	flag.BoolVar(&c.N, "n", false, "печатать номер строки")
	flag.Parse()

	if len(flag.Args()) < 1 {
		return errors.New("No pattern")
	} else {
		c.Pattern = flag.Arg(0)
		for i := 1; i < len(flag.Args()); i++ {
			c.Files = append(c.Files, flag.Arg(i))
		}
	}
	return nil
}
