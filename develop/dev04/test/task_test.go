package test

import (
	"lvl2/develop/dev04/internal/anagrams"
	"reflect"
	"testing"
)

func TestEqualWordsMap(t *testing.T) {
	tableTest := []struct {
		word1   map[rune]int
		word2   map[rune]int
		isEqual bool
	}{
		{
			map[rune]int{'о': 1, 'е': 1, 'н': 1, 'л': 1, 'с': 1, 'ц': 1},
			map[rune]int{'с': 1, 'о': 1, 'л': 1, 'н': 1, 'ц': 1, 'е': 1},
			true,
		},
		{
			map[rune]int{'о': 1, 'е': 1, 'н': 1, 'л': 1, 'с': 1, 'ц': 1},
			map[rune]int{'с': 1, 'о': 1, 'л': 1, 'н': 1, 'ц': 1, 'е': 1, 'а': 5},
			false,
		},
		{
			map[rune]int{'о': 1, 'е': 1, 'н': 1, 'л': 1, 'с': 1, 'ц': 1, 'a': 5},
			map[rune]int{'с': 1, 'о': 1, 'л': 1, 'н': 1, 'ц': 7, 'е': 1, 'а': 5},
			false,
		},
		{
			map[rune]int{'а': 1, 'и': 1, 'л': 1, 'б': 1, 'к': 1, 'е': 1, 'т': 2},
			map[rune]int{'т': 2, 'а': 1, 'б': 1, 'л': 1, 'е': 1, 'к': 1, 'и': 1},
			true,
		},
		{
			map[rune]int{'а': 100, 'и': 1, 'л': 4, 'б': 8, 'к': 1, 'е': 1, 'т': 2},
			map[rune]int{'т': 2, 'а': 100, 'б': 8, 'л': 4, 'е': 1, 'к': 1, 'и': 1},
			true,
		},
	}
	for _, val := range tableTest {
		if anagrams.EqualWordsMap(val.word1, val.word2) != val.isEqual {
			t.Errorf("Incorret compare: %v %v", val.word1, val.word2)
		}
	}
}

func TestMapOfWord(t *testing.T) {
	tableTest := []struct {
		word    string
		mapWord map[rune]int
	}{
		{
			"солнце",
			map[rune]int{'с': 1, 'о': 1, 'л': 1, 'н': 1, 'ц': 1, 'е': 1},
		},
		{
			"таблетки",
			map[rune]int{'т': 2, 'а': 1, 'б': 1, 'л': 1, 'е': 1, 'к': 1, 'и': 1},
		},
	}

	for _, val := range tableTest {
		if !anagrams.EqualWordsMap(anagrams.MapOfWord(val.word), val.mapWord) {
			t.Errorf("Incorret convertion word to map: %v", val.word)
		}
	}

}

func TestFindAnagrams(t *testing.T) {
	var words = []string{"пятАК", "пяТак", "Листок", "природа", "Пятка", "пятка", "стОлик",
		"Тяпка", "слиток"}
	anagramsMapTest := map[string]*[]string{
		"пятак":  {"пятка", "тяпка"},
		"листок": {"слиток", "столик"},
	}
	anagramsMapExe := anagrams.FindAnagrams(&words)
	if !reflect.DeepEqual(anagramsMapTest, anagramsMapExe) {
		t.Errorf("Incorret anagrams group")
	}
}
