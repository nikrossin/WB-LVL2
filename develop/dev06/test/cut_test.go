package test

import (
	"bytes"
	cut "lvl2/develop/dev06/internal"
	"reflect"
	"strings"
	"testing"
)

func TestFormatLine(t *testing.T) {
	c := &cut.Config{}

	Tests := []struct {
		line        string
		correctLine string
		D           string
		Fvalues     cut.FValues
	}{
		{
			"AA\tbb",
			"AA",
			"\t",
			cut.FValues{0: false},
		},
		{
			"AA\tbb\tcc\tfff\thhh",
			"AA\tbb\tcc",
			"\t",
			cut.FValues{0: false, 1: false, 2: false},
		},

		{
			"AA bb cc fff hhh 666 888 777 111 666",
			"bb fff hhh 666 888 777",
			" ",
			cut.FValues{1: false, 3: false, 4: false, 5: false, 6: false, 7: false},
		},
		{
			"AA bb cc fff hhh 666 888 777 111 666",
			"bb fff hhh 666 888 111 666",
			" ",
			cut.FValues{1: false, 3: false, 4: false, 5: false, 6: false, 8: true},
		},
	}

	for num, val := range Tests {
		c.D = val.D
		c.FValues = val.Fvalues
		cu := &cut.Cut{Config: c}
		newLine, err := cu.FormatLine(val.line)
		if err != nil {
			t.Errorf("Incorrect Error with FormatLine")
		}
		if newLine != val.correctLine {
			t.Errorf("Error with FormatLine %v %v %v", num, newLine, val.correctLine)
		}
	}

}

func TestRun(t *testing.T) {
	var in, out bytes.Buffer
	cfg := &cut.Config{}
	cfg.S = true
	cfg.D = " "
	cfg.F = "-2,4-5,7"
	cfg.FValues = make(cut.FValues)
	cfg.ParseFlagF()

	c := cut.NewCut(&in, &out)
	c.Config = cfg

	inputLine := []string{"aa bbbbbbbbbbbbbbbbbbbb cc dd ff ggggg hhh", "aaaa bbb", "aa\tbb"}
	correctLine := []string{"aa bbbbbbbbbbbbbbbbbbbb dd ff hhh", "aaaa bbb", ""}

	for _, v := range inputLine {
		in.WriteString(v + "\n")
	}
	if err := c.Run(); err != nil {
		t.Errorf("Incorrect error")
	}

	outData := strings.Split(out.String(), "\n")

	if !reflect.DeepEqual(outData, correctLine) {
		t.Errorf("Incorrect run with format lines %v %v", outData, correctLine)
	}

}
