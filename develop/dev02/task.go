package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func unpack(s string) (string, error) {
	var buffSymb rune
	var screen bool     // Включено ли экранирование
	var prevLetter bool // Предыдущий символ буква или нет
	var buffString strings.Builder

	for _, val := range s {
		if unicode.IsDigit(val) { // если текущий символ число
			if prevLetter { // если пред символ буква
				digit, err := strconv.Atoi(string(val))
				if err != nil {
					return "", err
				}
				buffString.WriteString(strings.Repeat(string(buffSymb), digit)) // добавляем в буффер букву n раз
				prevLetter = false                                              // пред символ цифра
			} else if screen { // если экранирование включено
				prevLetter = true // то на след итерации предыдущий символ считаем как букву
				buffSymb = val    // запоминаем текущий символ в буффер для дальнейшей записи в строку
				screen = false    // экранирование отключаем
			} else {
				return "", fmt.Errorf("Incorrect string") // иначе пред симвл тоже число - ошибка
			}
		} else if unicode.IsLetter(val) { // если текущий символ буква
			if prevLetter { // и предыдущий буква
				buffString.WriteRune(buffSymb) // то предыдущий символ записываем в буфер 1 раз
				buffSymb = val                 // запоминаем текущий символ в буффер для дальнейшей записи в строку
			} else {
				prevLetter = true
				buffSymb = val
			}

		} else if val == '\\' { // если экран
			if screen { // если экранирование уже включено
				prevLetter = true // то этот символ будет буквой
				buffSymb = val
				screen = false // выключаем экран
			} else if prevLetter { // если пред симол буква запишем его 1 раз в буфер
				buffString.WriteRune(buffSymb)
				prevLetter = false
				screen = true
			} else { // если цифра то включаем экранирование
				screen = true
			}

		}
	}
	if prevLetter {
		buffString.WriteRune(buffSymb)
	} else if screen && s != "" {
		return "", fmt.Errorf("Incorrect string")
	}
	return buffString.String(), nil
}

func main() {
	fmt.Println(unpack(`b5b6\`))
}
