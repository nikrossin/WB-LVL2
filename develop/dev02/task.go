package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func unpack(s string) (string, error) {
	var buffSymb rune
	var screen bool
	var prevLetter bool
	var buffString strings.Builder

	for _, val := range s {
		if unicode.IsDigit(val) {
			if prevLetter {
				digit, err := strconv.Atoi(string(val))
				if err != nil {
					return "", err
				}
				buffString.WriteString(strings.Repeat(string(buffSymb), digit))
				prevLetter = false
			} else if screen {
				prevLetter = true
				buffSymb = val
				screen = false
			} else {
				return "", fmt.Errorf("Incorrect string")
			}
		} else if unicode.IsLetter(val) {
			if prevLetter {
				buffString.WriteRune(buffSymb)
				buffSymb = val
			} else {
				prevLetter = true
				buffSymb = val
			}

		} else if val == '\\' {
			if screen {
				prevLetter = true
				buffSymb = val
				screen = false
			} else if prevLetter {
				buffString.WriteRune(buffSymb)
				prevLetter = false
				screen = true
			} else {
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
