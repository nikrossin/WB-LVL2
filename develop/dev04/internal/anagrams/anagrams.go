package anagrams

import (
	"sort"
	"strings"
)

type GroupAnagram struct {
	header    string
	headerMap map[rune]int
	words     []string
}

func NewGroupAnagram(header string, headerMap map[rune]int) *GroupAnagram {
	return &GroupAnagram{
		header:    header,
		headerMap: headerMap,
	}
}

func (g *GroupAnagram) AddWord(word string) bool {
	if word != g.header {
		for _, val := range g.words {
			if val == word {
				return false
			}
		}
	} else {
		return false
	}
	g.words = append(g.words, word)
	return true
}

func (g *GroupAnagram) SortWords() {
	sort.Strings(g.words)
}

type AnagramList struct {
	data []*GroupAnagram
}

func NewAnagramList() *AnagramList {
	return &AnagramList{
		data: []*GroupAnagram{},
	}
}

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

func wordToLower(word string) string {
	return strings.ToLower(word)
}

func (l *AnagramList) AddGroup(group *GroupAnagram) {
	l.data = append(l.data, group)
}
func (l *AnagramList) AddWordToList(word string) {
	isFindGroup := false
	wordMap := MapOfWord(word)
	for _, group := range l.data {
		if EqualWordsMap(group.headerMap, wordMap) {
			isFindGroup = group.AddWord(word)
			break
		}
	}
	if !isFindGroup {
		newGroup := NewGroupAnagram(word, wordMap)
		l.AddGroup(newGroup)
	}
}

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

func (l *AnagramList) GetFormatMap() map[string]*[]string {
	mapAnagrams := make(map[string]*[]string)
	for _, group := range l.data {
		mapAnagrams[group.header] = &group.words
	}
	return mapAnagrams
}

func FindAnagrams(words *[]string) map[string]*[]string {
	anagramList := NewAnagramList()
	for _, word := range *words {
		anagramList.AddWordToList(wordToLower(word))
	}
	anagramList = anagramList.SortWordsInGroupsAndClear()
	return anagramList.GetFormatMap()
}
