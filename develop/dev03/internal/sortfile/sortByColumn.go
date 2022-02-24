package sortfile

import (
	"sort"
)

func sortByColumn(input []string, flags *Flags) []string {
	resultData := make([]string, 0, len(input))
	mapData := make(map[string][]string)
	keys := make([]string, 0, len(input))

	for _, line := range input {
		str := parseString(line, flags)
		var column string // если колонки нет значение ""
		if flags.K <= len(str) {
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

	for _, key := range keys {
		for _, field := range mapData[key] {
			resultData = append(resultData, field)
		}
	}
	return resultData
}
