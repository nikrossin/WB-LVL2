package sortfile

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Flags struct {
	K int
	R bool
	U bool
	N bool
	M bool
	B bool
	C bool
	H bool
}

func NewFlags() *Flags {
	return &Flags{}
}
func (flags *Flags) ParseFlags() {
	flag.IntVar(&flags.K, "k", 1, "указание колонки для сортировки")
	flag.BoolVar(&flags.N, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&flags.R, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&flags.U, "u", false, "не выводить повторяющиеся строки")
	flag.BoolVar(&flags.M, "M", false, "сортировать по названию месяца")
	flag.BoolVar(&flags.B, "b", false, "игнорировать хвостовые пробелы")
	flag.BoolVar(&flags.C, "c", false, "проверять отсортированы ли данные")
	flag.BoolVar(&flags.H, "h", false, "сортировать по числовому значению с учётом суффиксов")
	flag.Parse()
}

type TextFile struct {
	*Flags
	dataStrings []string
	inputPath   string
	outputPath  string
	isSorted    bool
}

func NewTextFile() *TextFile {
	return &TextFile{
		Flags: NewFlags(),
	}
}

func (f *TextFile) Read() {
	file, err := os.Open(f.inputPath)

	if err != nil {
		fmt.Println("error opening file: err:", err)
		os.Exit(1)
	}
	defer file.Close()

	dataBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	f.dataStrings = strings.Split(string(dataBytes), "\n")

}

func (f *TextFile) Write() {
	if !f.C {
		file, err := os.Create(f.outputPath)

		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		for _, line := range f.dataStrings {
			file.WriteString(line + "\n")
		}
	} else {
		fmt.Println(f.isSorted)
	}

}

func (f *TextFile) SetInputPath() {
	f.inputPath = "../cmd/" + flag.Arg(0)

}

func (f *TextFile) SetOutputPath() {
	f.outputPath = "../cmd/" + flag.Arg(1)

}
