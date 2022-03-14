package cut

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"
)

// Cut Структура cut
type Cut struct {
	*Config
	input  io.Reader
	output io.Writer
}

// NewCut Создание нового cut
func NewCut(in io.Reader, out io.Writer) *Cut {
	return &Cut{input: in, output: out}
}

// InitConfig Инициализация конфигурации по аргументам
func (c *Cut) InitConfig() {
	config := NewConfigInit()
	err := config.ParseFlagF()
	if err != nil {
		log.Fatalln(err)
	}
	c.Config = config

}

// Run Запуск Cut
func (c *Cut) Run() error {
	sc := bufio.NewScanner(c.input)
	for sc.Scan() {
		line := sc.Text()
		cutLine, err := c.FormatLine(line) // "обрезка" строки
		if err != nil {
			return err
		}
		//не выводить если строка не имеет заданных столбцов и если включен режим S
		if !(!strings.Contains(line, c.D) && c.S) && cutLine != "" { //
			if _, err := fmt.Fprintln(c.output, cutLine); err != nil {
				return err
			}
		}
	}
	if err := sc.Err(); err != nil {
		return err
	}
	return nil
}

// FormatLine Форматирование строки согласно заданным столбцам
func (c *Cut) FormatLine(line string) (string, error) {
	lineColumns := strings.Split(line, c.D)
	if len(lineColumns) == 1 { // если столбец один -> не было разделителя, сразу возвращаем результат
		return line, nil
	}
	var formatString strings.Builder // формируем новую строку
	for index, col := range lineColumns {
		if isEndLine, ok := c.FValues[index]; ok {
			if _, err := formatString.WriteString(col + c.D); err != nil {
				return "", err
			}              // если столбец в флаге f есть, добавляем столбец к новой строке
			if isEndLine { // если значение столбца флага -f true, то добавляем все следующие столбцы к строке до конца
				for i := index + 1; i < len(lineColumns); i++ {
					if _, err := formatString.WriteString(lineColumns[i] + c.D); err != nil {
						return "", err
					}
				}
				newLine := strings.TrimSuffix(formatString.String(), c.D) // обрезаем конечный разделитель
				return newLine, nil
			}
		}

	}
	newLine := strings.TrimSuffix(formatString.String(), c.D)
	return newLine, nil
}
