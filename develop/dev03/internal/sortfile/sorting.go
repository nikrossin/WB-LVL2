package sortfile

import (
	"strings"
)

//отсечение хвостовых пробелов
func parseString(s string, flags *Flags) []string {
	if flags.B {
		return strings.Fields(s)
	} else {
		return strings.Split(s, " ")
	}
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func SortFile(file *TextFile) {
	//f.c втсроить в функции + f.b для basic sort
	var output []string
	if file.U {
		file.dataStrings = deleteDuplicateLines(file.dataStrings)
	}
	if file.M {
		output = SortByMonth(file.dataStrings, file.Flags)
	} else if file.H {
		output = SortBySuffix(file.dataStrings, file.Flags)
	} else if file.N {
		output = sortByNumbers(file.dataStrings, file.Flags)
	} else if file.K > 1 {
		output = sortByColumn(file.dataStrings, file.Flags)
	} else {
		output = basicSort(file.dataStrings, file.Flags)
	}

	if file.C {
		if equal(file.dataStrings, output) {
			file.isSorted = true
		} else {
			file.isSorted = false
		}
	} else {
		file.dataStrings = output
	}
}
