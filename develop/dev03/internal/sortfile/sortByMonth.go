package sortfile

import (
	"sort"
	"strings"
)

type Months []string

var month = map[string]int{
	"JAN": 1,
	"FAB": 2,
	"MAR": 3,
	"APR": 4,
	"MAY": 5,
	"JUN": 6,
	"JUL": 7,
	"AUG": 8,
	"SEP": 9,
	"OCT": 10,
	"NOV": 11,
	"DEC": 12,
}

func (m Months) Len() int {
	return len(m)
}
func (m Months) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func (m Months) Less(i, j int) bool {
	return month[m[i]] < month[m[j]]
}

func SortByMonth(input []string, flags *Flags) []string {

	resultData := make([]string, 0, len(input))
	mapData := make(map[string][]string) //столбец сортировки:полные строки
	keys := make(Months, 0, len(input))  // столбцы для сортировки

	for _, line := range input {
		str := parseString(line, flags)
		var column string        // если колонки нет значение ""
		if flags.K <= len(str) { // если столбец существует в строке
			column = str[flags.K-1]
			if _, ok := month[strings.ToUpper(column[:3])]; ok {
				column = strings.ToUpper(column[:3])
			} else {
				column = ""
			}
		}
		if _, ok := mapData[column]; !ok {
			keys = append(keys, column)
		}
		mapData[column] = append(mapData[column], line)

	}
	//fmt.Println(keys)
	if flags.R {

		sort.Sort(sort.Reverse(keys))
	} else {
		sort.Sort(keys)
	}
	//fmt.Println(keys)
	for _, key := range keys {
		for _, field := range mapData[key] {
			resultData = append(resultData, field)
		}
	}
	return resultData
}
