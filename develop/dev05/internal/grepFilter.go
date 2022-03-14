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

// Filter Фильтр файла
type Filter struct {
	Source string
	Config *Config
	wg     *sync.WaitGroup
	Lines  []string
}

// NewFilter Новый фильтр
func NewFilter(s string, c *Config, w *sync.WaitGroup) *Filter {
	return &Filter{
		Config: c,
		Source: s,
		wg:     w,
	}
}

// RunFilter Запуск фильтрации
func (f *Filter) RunFilter() {
	defer f.wg.Done()
	if f.Source != "" { // фильтрация по файлу
		if err := f.FilterFile(os.Stdout); err != nil {
			log.Fatalln(err)
		}
	} else { // фильрация по stdin
		if err := f.FilterStdin(os.Stdout); err != nil {
			log.Fatalln(err)
		}
	}
}

// ReadFile Чтение из файла строк в Lines
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

// FilterFile Фильтрация данных файла
func (f *Filter) FilterFile(w io.Writer) error {
	if err := f.ReadFile(); err != nil {
		log.Fatal(err)
	}
	var count int
	for i, line := range f.Lines {

		if f.AccordanceConfig(line) { // проверка подходит ли строка по значениям флагов
			if f.Config.C { // Подсчет количества строк если флаг установлен
				count++
			} else if f.Config.AA > 0 { // -A
				if err := f.PrintLine(w, i); err != nil {
					return err
				}
				if err := f.PrintAfter(w, i, f.Config.AA); err != nil {
					return err
				}
			} else if f.Config.BB > 0 { // -B
				if err := f.PrintBefore(w, i, f.Config.BB); err != nil {
					return err
				}
				if err := f.PrintLine(w, i); err != nil {
					return err
				}
			} else if f.Config.CC > 0 { // -C
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
	if f.Config.C { // вывод количества строк
		if _, err := fmt.Fprintln(w, count); err != nil {
			return err
		}
	}
	return nil
}

// FilterStdin фильтрация данных из stdio
func (f *Filter) FilterStdin(w io.Writer) error {
	var countAfterLine, index int
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		line := sc.Text()
		f.Lines = append(f.Lines, line)
		if countAfterLine > 0 { // счетчик вывода строк с флагом -A ( т.к. ввод данных построчный)
			if err := f.PrintLine(w, index); err != nil {
				return err
			}
			countAfterLine--
		}

		if f.AccordanceConfig(line) { // если строка соотвествует параметрам флагов
			countAfterLine = 0 // обнуляем счетчик даже если не все строки вывелись полсе предыдущего соответсвия

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
		index++ // индекс строки
	}
	return nil
}

// AccordanceConfig Проверка соответствия строки Фильтру по флагам
func (f *Filter) AccordanceConfig(line string) bool {
	pattern := f.Config.Pattern

	if f.Config.F { // отключение regexp
		if f.Config.I {
			pattern = strings.ToLower(pattern)
			line = strings.ToLower(line)
		}
		if strings.Contains(line, pattern) {
			return !f.Config.V
		}
	} else {
		if f.Config.I {
			pattern = "(?i)" + pattern // regexp регистронезависимый
		}
		if val, _ := regexp.MatchString(pattern, line); val {
			return !f.Config.V
		}

	}
	return f.Config.V // флаг инвертирования соответсвия строки флагам
}

// PrintLine Вывод строки, соответсвующей значениям фильтра
func (f *Filter) PrintLine(w io.Writer, index int) error {
	path := "stdout" // префикс - откуда считана строка
	if f.Source != "" {
		path = f.Source
	}
	// проверка при флагах -A,B,C, что не вышли за диапазон строк
	if index < len(f.Lines) && index >= 0 {
		line := f.Lines[index]
		if f.Config.N { // если флаг нумерации строк задан
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

// PrintAfter Вывод строк при -A,C
func (f *Filter) PrintAfter(w io.Writer, start int, delta int) error {
	for i := start + 1; i <= start+delta; i++ {
		if err := f.PrintLine(w, i); err != nil {
			return err
		}
	}
	return nil
}

// PrintBefore Вывод строк при -B,C
func (f *Filter) PrintBefore(w io.Writer, end int, delta int) error {
	for i := end - delta; i < end; i++ {
		if err := f.PrintLine(w, i); err != nil {
			return err
		}
	}
	return nil
}
