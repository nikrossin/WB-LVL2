package sortfile

// Удаление строк дупликатов
func deleteDuplicateLines(input []string) []string {
	mapData := make(map[string]struct{}) // мапа с уникальными строками
	for _, line := range input {
		if _, ok := mapData[line]; !ok {
			mapData[line] = struct{}{}
		}
	}
	resultData := make([]string, 0, len(mapData))
	for key := range mapData {
		resultData = append(resultData, key)
	}
	return resultData
}
