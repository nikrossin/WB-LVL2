package sortfile

import (
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

var suffixes = map[string]int{"n": -9, "mi": -6, "m": -3, "K": 3, "M": 6, "G": 9}

// получить абсолютное число из числа с суффиксом
func getNum(s string) float64 {
	for suffix, degree := range suffixes {
		if strings.HasSuffix(s, suffix) {
			//если суффикс верный, отделяем число
			fStr, err := strconv.ParseFloat(strings.TrimSuffix(s, suffix), 32)
			//если число некорректрое, возвращаем -inf
			if err != nil {
				fStr, err = strconv.ParseFloat("-inf", 32)
				if err != nil {
					log.Fatalln(err)
				}
				return fStr
			}
			return math.Pow10(degree) * fStr
		}
	}
	//если суффик не найден, проверяем не "чистое" ли это число
	fStr, err := strconv.ParseFloat(s, 32)
	if err != nil {
		fStr, err = strconv.ParseFloat("-inf", 32) // если не число -inf
		if err != nil {
			log.Fatalln(err)
		}
	}
	return fStr
}

// SortBySuffix Сортировка по числовому суффиксу
func SortBySuffix(input []string, flags *Flags) []string {
	resultData := make([]string, 0, len(input))
	mapData := make(map[float64][]string)  //столбец сортировки:полные строки
	keys := make([]float64, 0, len(input)) // столбцы для сортировки

	for _, line := range input {
		str := parseString(line, flags)
		column, err := strconv.ParseFloat("-inf", 32) // если колонки нет значение -infinity
		if err != nil {
			log.Fatalln(err)
		}
		if flags.K <= len(str) { // если столбец существует в строке
			column = getNum(str[flags.K-1])
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
