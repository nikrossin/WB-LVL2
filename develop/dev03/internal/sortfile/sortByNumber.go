package sortfile

import (
	"log"
	"sort"
	"strconv"
)

// сортировка по числу столбца
func sortByNumbers(input []string, flags *Flags) []string {
	resultData := make([]string, 0, len(input))
	mapData := make(map[float64][]string)
	keys := make([]float64, 0, len(input))

	for _, line := range input {
		str := parseString(line, flags)
		column, err := strconv.ParseFloat("-inf", 32) // если колонки нет значение -infinity
		if err != nil {
			log.Fatalln(err)
		}
		if flags.K <= len(str) {
			column, err = strconv.ParseFloat(str[flags.K-1], 32) // парсим строку в число
			if err != nil {
				column, _ = strconv.ParseFloat("-inf", 32) // если не число, то приравниваем -inf для сортировки
				if err != nil {
					log.Fatalln(err)
				}
			}
		}
		if _, ok := mapData[column]; !ok {
			keys = append(keys, column)
		}
		mapData[column] = append(mapData[column], line)
	}
	if flags.R {
		sort.Sort(sort.Reverse(sort.Float64Slice(keys)))
	} else {
		sort.Float64s(keys)
	}

	for _, key := range keys {
		for _, field := range mapData[key] {
			resultData = append(resultData, field)
		}
	}
	return resultData
}
