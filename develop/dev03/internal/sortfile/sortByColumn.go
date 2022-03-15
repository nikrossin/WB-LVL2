package sortfile

import (
	"sort"
)

// сортировка по заданной колонке
func sortByColumn(input []string, flags *Flags) []string {
	resultData := make([]string, 0, len(input)) // результат сортировки
	mapData := make(map[string][]string)        // ключ - сортируемый столбец, значение - полная строка
	keys := make([]string, 0, len(input))       // сортируем ключи (столбцы)

	for _, line := range input {
		str := parseString(line, flags) // преобразование строки в колонки
		var column string               // если колонки нет значение ""
		if flags.K <= len(str) {        //  проверка что заданная колонка не превышает количества колонок строки
			column = str[flags.K-1]
		}
		if _, ok := mapData[column]; !ok {
			keys = append(keys, column)
		}
		mapData[column] = append(mapData[column], line)

	}
	if flags.R {
		sort.Sort(sort.Reverse(sort.StringSlice(keys)))
	} else {
		sort.Strings(keys)
	}

	for _, key := range keys { // результат сортировки
		for _, field := range mapData[key] {
			resultData = append(resultData, field)
		}
	}
	return resultData
}
