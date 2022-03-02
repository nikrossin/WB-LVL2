package grep

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"sync"
)

type Filter struct {
	Source string
	Config *Config
	wg     *sync.WaitGroup
	Lines  []string
}

func NewFilter(s string, c *Config, w *sync.WaitGroup) *Filter {
	return &Filter{
		Config: c,
		Source: s,
		wg:     w,
	}
}
func (f *Filter) RunFilter() {
	defer f.wg.Done()
	if f.Source != "" {
		if err := f.FilterFile(os.Stdout); err != nil {
			log.Fatalln(err)
		}
	} else {
		if err := f.FilterStdin(os.Stdout); err != nil {
			log.Fatalln(err)
		}
	}
}
func (f *Filter) ReadFile() error {
	file, err := os.Open(f.Source)

	if err != nil {
		return fmt.Errorf("error opening file: err: %v", err)

	}
	defer file.Close()

	dataBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("error read data from file: %v", err)
	}
	f.Lines = strings.Split(string(dataBytes), "\n")
	return nil
}
func (f *Filter) FilterFile(w io.Writer) error {
	if err := f.ReadFile(); err != nil {
		log.Fatal(err)
	}
	var count int
	for i, line := range f.Lines {

		if f.AccordanceConfig(line) {
			if f.Config.C {
				count++
			} else if f.Config.AA > 0 {
				if err := f.PrintLine(w, i); err != nil {
					return err
				}
				if err := f.PrintAfter(w, i, f.Config.AA); err != nil {
					return err
				}
			} else if f.Config.BB > 0 {
				if err := f.PrintBefore(w, i, f.Config.BB); err != nil {
					return err
				}
				if err := f.PrintLine(w, i); err != nil {
					return err
				}
			} else if f.Config.CC > 0 {
				if err := f.PrintBefore(w, i, f.Config.CC); err != nil {
					return err
				}
				if err := f.PrintLine(w, i); err != nil {
					return err
				}
				if err := f.PrintAfter(w, i, f.Config.CC); err != nil {
					return err
				}
			} else {
				if err := f.PrintLine(w, i); err != nil {
					return err
				}
			}

		}

	}
	if f.Config.C {
		if _, err := fmt.Fprintln(w, count); err != nil {
			return err
		}
	}
	return nil
}

func (f *Filter) FilterStdin(w io.Writer) error {
	var countAfterLine, index int
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		line := sc.Text()
		f.Lines = append(f.Lines, line)
		if countAfterLine > 0 {
			if err := f.PrintLine(w, index); err != nil {
				return err
			}
			countAfterLine--
		}

		if f.AccordanceConfig(line) {
			countAfterLine = 0

			if f.Config.AA > 0 {
				countAfterLine += f.Config.AA
				if err := f.PrintLine(w, index); err != nil {
					return err
				}
			} else if f.Config.BB > 0 {
				if err := f.PrintBefore(w, index, f.Config.BB); err != nil {
					return err
				}
				if err := f.PrintLine(w, index); err != nil {
					return err
				}
			} else if f.Config.CC > 0 {
				if err := f.PrintBefore(w, index, f.Config.CC); err != nil {
					return err
				}
				if err := f.PrintLine(w, index); err != nil {
					return err
				}
				countAfterLine += f.Config.CC
			} else {
				if err := f.PrintLine(w, index); err != nil {
					return err
				}
			}
		}
		index++
	}
	return nil
}

func (f *Filter) AccordanceConfig(line string) bool {
	pattern := f.Config.Pattern

	if f.Config.F {
		if f.Config.I {
			pattern = strings.ToLower(pattern)
			line = strings.ToLower(line)
		}
		if strings.Contains(line, pattern) {
			return !f.Config.V
		}
	} else {
		if f.Config.I {
			pattern = "(?i)" + pattern
		}
		if val, _ := regexp.MatchString(pattern, line); val {
			return !f.Config.V
		}

	}
	return f.Config.V
}

func (f *Filter) PrintLine(w io.Writer, index int) error {
	path := "stdout"
	if f.Source != "" {
		path = f.Source
	}
	if index < len(f.Lines) && index >= 0 {
		line := f.Lines[index]
		if f.Config.N {
			line = fmt.Sprintf("%v: %v:%v", path, index+1, line)
		} else {
			line = fmt.Sprintf("%v: %v", path, line)
		}
		if _, err := fmt.Fprintln(w, line); err != nil {
			return err
		}
	}
	return nil
}
func (f *Filter) PrintAfter(w io.Writer, start int, delta int) error {
	for i := start + 1; i <= start+delta; i++ {
		if err := f.PrintLine(w, i); err != nil {
			return err
		}
	}
	return nil
}

func (f *Filter) PrintBefore(w io.Writer, end int, delta int) error {
	for i := end - delta; i < end; i++ {
		if err := f.PrintLine(w, i); err != nil {
			return err
		}
	}
	return nil
}
