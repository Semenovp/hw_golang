package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	var sb strings.Builder
	var prevValue rune
	var prevIsLetter bool
	var currentValue rune
	var currentIsLetter bool
	for i, val := range input {
		currentValue = val
		currentIsLetter = unicode.IsLetter(val)

		if unicode.IsDigit(val) {
			if prevIsLetter {
				sb.WriteString(strings.Repeat(string(prevValue), int(currentValue-'0')))
			} else {
				return "", ErrInvalidString
			}
		} else if prevIsLetter {
			sb.WriteRune(prevValue)
		}

		if currentIsLetter && i == len(input)-1 {
			sb.WriteRune(currentValue)
		}

		prevValue = currentValue
		prevIsLetter = currentIsLetter
	}

	return sb.String(), nil
}
