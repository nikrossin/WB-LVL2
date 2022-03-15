package sortfile

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Flags Структура хранящая все флаги
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

// NewFlags Создание Flags
func NewFlags() *Flags {
	return &Flags{}
}

// ParseFlags Инициализация стуктуры флагами
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

// TextFile Структура с конфигурацией файла для сортировки
type TextFile struct {
	*Flags
	dataStrings []string
	inputPath   string
	outputPath  string
	isSorted    bool
}

// NewTextFile Создание TextFile
func NewTextFile() *TextFile {
	return &TextFile{
		Flags: NewFlags(),
	}
}

// Чтение из файла и запись в dataString
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

// Запись отсортированных строк в фаил
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

// SetInputPath Установка пути файла для чтения
func (f *TextFile) SetInputPath() {
	f.inputPath = flag.Arg(0)

}

// SetOutputPath Установка пути файла для записи
func (f *TextFile) SetOutputPath() {
	f.outputPath = flag.Arg(1)

}
