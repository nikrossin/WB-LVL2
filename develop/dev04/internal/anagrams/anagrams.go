package anagrams

import (
	"sort"
	"strings"
)

// GroupAnagram Структура группы слов анаграмм
type GroupAnagram struct {
	header    string
	headerMap map[rune]int
	words     []string
}

// NewGroupAnagram Создать новую группу анаграмм
func NewGroupAnagram(header string, headerMap map[rune]int) *GroupAnagram {
	return &GroupAnagram{
		header:    header,
		headerMap: headerMap,
	}
}

// AddWord Добавление в группу анаграмм слова, если его еще нет в группе
func (g *GroupAnagram) AddWord(word string) {
	if word != g.header {
		for _, val := range g.words {
			if val == word {
				return
			}
		}
	} else {
		return
	}
	g.words = append(g.words, word)
}

// SortWords Сортировка слов в группе
func (g *GroupAnagram) SortWords() {
	sort.Strings(g.words)
}

// AnagramList Список групп анаграмм
type AnagramList struct {
	data []*GroupAnagram
}

// NewAnagramList Создать новый список
func NewAnagramList() *AnagramList {
	return &AnagramList{
		data: []*GroupAnagram{},
	}
}

// MapOfWord создание хеш-таблицы символов слова
func MapOfWord(word string) map[rune]int {
	mapWord := make(map[rune]int)
	for _, symb := range word {
		if _, ok := mapWord[symb]; ok {
			mapWord[symb]++
		} else {
			mapWord[symb] = 1
		}
	}
	return mapWord
}

// EqualWordsMap сравнение на равенство 2-х хеш-таблиц
func EqualWordsMap(word1 map[rune]int, word2 map[rune]int) bool {
	if len(word1) != len(word2) {
		return false
	}
	for symb, count := range word1 {
		if count2, ok := word2[symb]; ok {
			if count != count2 {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

// Привести слово в нижний регистр
func wordToLower(word string) string {
	return strings.ToLower(word)
}

// AddGroup Добавляет группу анаграмм в список
func (l *AnagramList) AddGroup(group *GroupAnagram) {
	l.data = append(l.data, group)
}

// AddWordToList Добавляет новове слово в соответствующую группу, если такой нет, создает группу
func (l *AnagramList) AddWordToList(word string) {

	wordMap := MapOfWord(word)
	for _, group := range l.data {
		if EqualWordsMap(group.headerMap, wordMap) {
			group.AddWord(word)
			return
		}
	}
	newGroup := NewGroupAnagram(word, wordMap)
	l.AddGroup(newGroup)

}

// SortWordsInGroupsAndClear Удаляет группы только с одним словом и сортирует слова в группе
func (l *AnagramList) SortWordsInGroupsAndClear() *AnagramList {
	anagramList := NewAnagramList()
	for _, group := range l.data {
		if len(group.words) != 0 {
			group.SortWords()
			anagramList.data = append(anagramList.data, group)
		}
	}
	return anagramList
}

// GetFormatMap Создает из списка групп мапу с ключем(первое словов) и указателем на слайс слов группы
func (l *AnagramList) GetFormatMap() map[string]*[]string {
	mapAnagrams := make(map[string]*[]string)
	for _, group := range l.data {
		mapAnagrams[group.header] = &group.words
	}
	return mapAnagrams
}

// FindAnagrams Поиск Анаграмм по слайсу слов (функция запуска)
func FindAnagrams(words *[]string) map[string]*[]string {
	anagramList := NewAnagramList()
	for _, word := range *words {
		anagramList.AddWordToList(wordToLower(word))
	}
	anagramList = anagramList.SortWordsInGroupsAndClear()
	return anagramList.GetFormatMap()
}
