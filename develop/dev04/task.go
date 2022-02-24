package main

import (
	"fmt"
	"lvl2/develop/dev04/internal/anagrams"
)

func main() {
	var words = []string{"пятак", "пятак", "листок", "природа", "пятка", "пятка", "столик",
		"тяпка", "слиток"}
	m := anagrams.FindAnagrams(&words)
	for key, _ := range m {
		fmt.Println(key, *m[key])
	}
}
