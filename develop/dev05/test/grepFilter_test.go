package test

import (
	"bytes"
	grep "lvl2/develop/dev05/internal"
	"strings"
	"testing"
)

func TestAccordanceConfig(t *testing.T) {

	lines := []string{"abc", "AbC", "abcAAA"}
	c1 := grep.NewConfig()
	c1.F = true
	c1.Pattern = "abc"
	f1 := grep.Filter{Config: c1}

	c2 := grep.NewConfig()
	c2.V = true
	c2.Pattern = "abc"
	f2 := grep.Filter{Config: c2}

	c3 := grep.NewConfig()
	c3.I = true
	c3.Pattern = "abc"
	f3 := grep.Filter{Config: c3}

	c4 := grep.NewConfig()
	c4.I, c4.V, c4.F = true, true, true
	c4.Pattern = "abc"
	f4 := grep.Filter{Config: c4}

	if f1.AccordanceConfig(lines[0]) != true || f1.AccordanceConfig(lines[1]) != false || f1.AccordanceConfig(lines[2]) != true {
		t.Errorf("Inccorect check with flag f\n")
	}
	if f2.AccordanceConfig(lines[0]) != false || f2.AccordanceConfig(lines[1]) != true || f2.AccordanceConfig(lines[2]) != false {
		t.Errorf("Inccorect check with flag v\n")
	}
	if f3.AccordanceConfig(lines[0]) != true || f3.AccordanceConfig(lines[1]) != true || f3.AccordanceConfig(lines[2]) != true {
		t.Errorf("Inccorect check with flag i\n")
	}
	if f4.AccordanceConfig(lines[0]) != false || f4.AccordanceConfig(lines[1]) != false || f4.AccordanceConfig(lines[2]) != false {
		t.Errorf("Inccorect check with flag v&i&f\n")
	}
}

func TestPrintLine(t *testing.T) {
	var buff bytes.Buffer
	c := grep.NewConfig()
	f := grep.Filter{Config: c, Lines: []string{"ABC"}}

	err1 := f.PrintLine(&buff, 0)
	r1 := buff.String()
	buff.Reset()

	f.Source = "file.txt"
	err2 := f.PrintLine(&buff, 0)
	r2 := buff.String()
	buff.Reset()

	err3 := f.PrintLine(&buff, 2)
	r3 := buff.String()
	buff.Reset()

	c.N = true
	err4 := f.PrintLine(&buff, 0)
	r4 := buff.String()
	buff.Reset()

	exp1 := "stdout: ABC\n"
	exp2 := "file.txt: ABC\n"
	exp3 := ""
	exp4 := "file.txt: 1:ABC\n"

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		t.Errorf("error while works PrintLine\n")
	}
	if exp1 != r1 {
		t.Errorf("incorrect output, expected: %s, got: %s\n", exp1, r1)
	}
	if exp2 != r2 {
		t.Errorf("incorrect output, expected: %s, got: %s\n", exp2, r2)
	}
	if exp3 != r3 {
		t.Errorf("incorrect output, expected: %s, got: %s\n", exp3, r3)
	}
	if exp4 != r4 {
		t.Errorf("incorrect output, expected: %s, got: %s\n", exp4, r4)
	}

}

func TestFilterFileA(t *testing.T) {
	var buff bytes.Buffer
	c := grep.NewConfig()
	s := grep.Filter{
		Config: c,
	}
	s.Source = "file.txt"
	exp := []string{"file.txt: abc", "file.txt: abCcc", "file.txt: AbCCCCCCC", "file.txt: Hello", "file.txt: sssabcaa", ""}

	c.AA = 3
	c.Pattern = "abc"
	err := s.FilterFile(&buff)
	res := strings.Split(buff.String(), "\n")

	if err != nil {
		t.Errorf("incorrect work of FilterFile with A flag\n")
	}
	if len(res) != len(exp) {
		t.Errorf("incorrect length of output resualt %v\n%v", exp, res)
	} else {
		for i := range exp {
			if res[i] != exp[i] {
				t.Errorf("incorrect output with flag -A, expected: %s, got: %s\n", exp[i], res[i])
			}
		}
	}
}

func TestFilterFileB(t *testing.T) {
	var buff bytes.Buffer
	c := grep.NewConfig()
	s := grep.Filter{
		Config: c,
	}
	s.Source = "file.txt"
	exp := []string{"file.txt: abc", "file.txt: Print", "file.txt: ghjk", "file.txt: ssss", "file.txt: sssabcaa", ""}

	c.BB = 3
	c.Pattern = "abc"
	err := s.FilterFile(&buff)
	res := strings.Split(buff.String(), "\n")

	if err != nil {
		t.Errorf("incorrect work of FilterFile with B flag\n")
	}
	if len(res) != len(exp) {
		t.Errorf("incorrect length of output resualt %v\n%v", exp, res)
	} else {
		for i := range exp {
			if res[i] != exp[i] {
				t.Errorf("incorrect output with flag B, expected: %s, got: %s\n", exp[i], res[i])
			}
		}
	}
}

func TestFilterFileC(t *testing.T) {
	var buff bytes.Buffer
	c := grep.NewConfig()
	s := grep.Filter{
		Config: c,
	}
	s.Source = "file.txt"
	exp := []string{"file.txt: abc", "file.txt: abCcc", "file.txt: AbCCCCCCC", "file.txt: ghjk", "file.txt: ssss",
		"file.txt: sssabcaa", ""}

	c.CC = 2
	c.Pattern = "abc"
	err := s.FilterFile(&buff)
	res := strings.Split(buff.String(), "\n")

	if err != nil {
		t.Errorf("incorrect work of FilterFile with C flag\n")
	}
	if len(res) != len(exp) {
		t.Errorf("incorrect length of output resualt %v\n%v", exp, res)
	} else {
		for i := range exp {
			if res[i] != exp[i] {
				t.Errorf("incorrect output with flag C, expected: %s, got: %s\n", exp[i], res[i])
			}
		}
	}
}

func TestFilterFileCI(t *testing.T) {
	var buff bytes.Buffer
	c := grep.NewConfig()
	s := grep.Filter{
		Config: c,
	}
	s.Source = "file.txt"
	exp := []string{"file.txt: abCcc", "file.txt: AbCCCCCCC", "file.txt: Hello", "file.txt: Print", "file.txt: ghjk", ""}

	c.CC = 2
	c.I = true
	c.Pattern = "hel"
	err := s.FilterFile(&buff)
	res := strings.Split(buff.String(), "\n")

	if err != nil {
		t.Errorf("incorrect work of FilterFile with CI flag\n")
	}
	if len(res) != len(exp) {
		t.Errorf("incorrect length of output resualt %v\n%v", exp, res)
	} else {
		for i := range exp {
			if res[i] != exp[i] {
				t.Errorf("incorrect output with flag CI, expected: %s, got: %s\n", exp[i], res[i])
			}
		}
	}
}

func TestFilterFileVI(t *testing.T) {
	var buff bytes.Buffer
	c := grep.NewConfig()
	s := grep.Filter{
		Config: c,
	}
	s.Source = "file.txt"
	exp := []string{"file.txt: Hello", "file.txt: Print", "file.txt: ghjk", "file.txt: ssss", ""}

	c.V = true
	c.I = true
	c.Pattern = "abc"
	err := s.FilterFile(&buff)
	res := strings.Split(buff.String(), "\n")

	if err != nil {
		t.Errorf("incorrect work of FilterFile with VI flag\n")
	}
	if len(res) != len(exp) {
		t.Errorf("incorrect length of output resualt %v\n%v", exp, res)
	} else {
		for i := range exp {
			if res[i] != exp[i] {
				t.Errorf("incorrect output with flag VI, expected: %s, got: %s\n", exp[i], res[i])
			}
		}
	}
}

func TestFilterFileBIN(t *testing.T) {
	var buff bytes.Buffer
	c := grep.NewConfig()
	s := grep.Filter{
		Config: c,
	}
	s.Source = "file.txt"
	exp := []string{"file.txt: 3:AbCCCCCCC", "file.txt: 4:Hello", ""}

	c.BB = 1
	c.I = true
	c.N = true
	c.Pattern = "hel"
	err := s.FilterFile(&buff)
	res := strings.Split(buff.String(), "\n")

	if err != nil {
		t.Errorf("incorrect work of FilterFile with BIN flag\n")
	}
	if len(res) != len(exp) {
		t.Errorf("incorrect length of output resualt %v\n%v", exp, res)
	} else {
		for i := range exp {
			if res[i] != exp[i] {
				t.Errorf("incorrect output with flag BIN, expected: %s, got: %s\n", exp[i], res[i])
			}
		}
	}
}

func TestFilterFileIc(t *testing.T) {
	var buff bytes.Buffer
	c := grep.NewConfig()
	s := grep.Filter{
		Config: c,
	}
	s.Source = "file.txt"
	exp := []string{"4", ""}

	c.I = true
	c.C = true
	c.Pattern = "abc"
	err := s.FilterFile(&buff)
	res := strings.Split(buff.String(), "\n")

	if err != nil {
		t.Errorf("incorrect work of FilterFile with Ic flag\n")
	}
	if len(res) != len(exp) {
		t.Errorf("incorrect length of output resualt %v\n%v", exp, res)
	} else {
		for i := range exp {
			if res[i] != exp[i] {
				t.Errorf("incorrect output with flag Ic, expected: %s, got: %s\n", exp[i], res[i])
			}
		}
	}
}

func TestFilterFileDefault(t *testing.T) {
	var buff bytes.Buffer
	c := grep.NewConfig()
	s := grep.Filter{
		Config: c,
	}
	s.Source = "file.txt"
	exp := []string{"file.txt: abc", ""}

	c.Pattern = "^abc"
	err := s.FilterFile(&buff)
	res := strings.Split(buff.String(), "\n")

	if err != nil {
		t.Errorf("incorrect work of FilterFile with none flag\n")
	}
	if len(res) != len(exp) {
		t.Errorf("incorrect length of output resualt %v\n%v", exp, res)
	} else {
		for i := range exp {
			if res[i] != exp[i] {
				t.Errorf("incorrect output with flag none, expected: %s, got: %s\n", exp[i], res[i])
			}
		}
	}
}

func TestFilterFileF1(t *testing.T) {
	var buff bytes.Buffer
	c := grep.NewConfig()
	s := grep.Filter{
		Config: c,
	}
	s.Source = "file.txt"
	exp := []string{""}

	c.F = true
	c.Pattern = "^abc"
	err := s.FilterFile(&buff)
	res := strings.Split(buff.String(), "\n")

	if err != nil {
		t.Errorf("incorrect work of FilterFile with F(1) flag\n")
	}
	if len(res) != len(exp) {
		t.Errorf("incorrect length of output resualt %v\n%v", exp, res)
	} else {
		for i := range exp {
			if res[i] != exp[i] {
				t.Errorf("incorrect output with flag F(1), expected: %s, got: %s\n", exp[i], res[i])
			}
		}
	}
}

func TestFilterFileF2(t *testing.T) {
	var buff bytes.Buffer
	c := grep.NewConfig()
	s := grep.Filter{
		Config: c,
	}
	s.Source = "file.txt"
	exp := []string{"file.txt: abc", "file.txt: sssabcaa", ""}

	c.F = true
	c.Pattern = "abc"
	err := s.FilterFile(&buff)
	res := strings.Split(buff.String(), "\n")

	if err != nil {
		t.Errorf("incorrect work of FilterFile with F(2) flag\n")
	}
	if len(res) != len(exp) {
		t.Errorf("incorrect length of output resualt %v\n%v", exp, res)
	} else {
		for i := range exp {
			if res[i] != exp[i] {
				t.Errorf("incorrect output with flag F(3), expected: %s, got: %s\n", exp[i], res[i])
			}
		}
	}
}
