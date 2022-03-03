package cut

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"
)

type Cut struct {
	*Config
	input  io.Reader
	output io.Writer
}

func NewCut(in io.Reader, out io.Writer) *Cut {
	return &Cut{input: in, output: out}
}
func (c *Cut) InitConfig() {
	config := NewConfigInit()
	err := config.ParseFlagF()
	if err != nil {
		log.Fatalln(err)
	}
	c.Config = config

}

func (c *Cut) Run() error {
	sc := bufio.NewScanner(c.input)
	for sc.Scan() {
		line := sc.Text()
		cutLine, err := c.FormatLine(line)
		if err != nil {
			return err
		}
		if !(!strings.Contains(line, c.D) && c.S) && cutLine != "" {
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

func (c *Cut) FormatLine(line string) (string, error) {
	lineColumns := strings.Split(line, c.D)
	if len(lineColumns) == 1 {
		return line, nil
	}
	var formatString strings.Builder
	for index, col := range lineColumns {
		if isEndLine, ok := c.FValues[index]; ok {
			if _, err := formatString.WriteString(col + c.D); err != nil {
				return "", err
			}
			if isEndLine {
				for i := index + 1; i < len(lineColumns); i++ {
					if _, err := formatString.WriteString(lineColumns[i] + c.D); err != nil {
						return "", err
					}
				}
				newLine := strings.TrimSuffix(formatString.String(), c.D)
				return newLine, nil
			}
		}

	}
	newLine := strings.TrimSuffix(formatString.String(), c.D)
	return newLine, nil
}
