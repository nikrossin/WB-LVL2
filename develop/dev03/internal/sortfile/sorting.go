package sortfile

import (
	"strings"
)

//отсечение хвостовых пробелов
func parseString(s string, flags *Flags) []string {
	if flags.B {
		return strings.Fields(s) // осекаем все пробелы и разбиваем строку на слайс
	}
	return strings.Split(s, " ") // разбиваем по пробелу одному
}

// сравнение на равенство 2-х слайсов
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

// SortFile Функция запуска сортировки
func SortFile(file *TextFile) {

	var output []string
	if file.U { // Удаление дублирования
		file.dataStrings = deleteDuplicateLines(file.dataStrings)
	}
	if file.M { // по месяцам
		output = SortByMonth(file.dataStrings, file.Flags)
	} else if file.H { // по суффиксам
		output = SortBySuffix(file.dataStrings, file.Flags)
	} else if file.N { // по месяцам
		output = sortByNumbers(file.dataStrings, file.Flags)
	} else if file.K > 1 { // по столбцу
		output = sortByColumn(file.dataStrings, file.Flags)
	} else { // базовая сортировка
		output = basicSort(file.dataStrings, file.Flags)
	}

	if file.C { // проверка отсортированы ли данные
		if equal(file.dataStrings, output) {
			file.isSorted = true
		} else {
			file.isSorted = false
		}
	} else {
		file.dataStrings = output
	}
}
