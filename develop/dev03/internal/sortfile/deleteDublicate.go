package sortfile

func deleteDuplicateLines(input []string) []string {
	mapData := make(map[string]struct{})
	for _, line := range input {
		if _, ok := mapData[line]; !ok {
			mapData[line] = struct{}{}
		}
	}
	resultData := make([]string, 0, len(mapData))
	for key, _ := range mapData {
		resultData = append(resultData, key)
	}
	return resultData
}
