package rle

import (
	"errors"
	"strconv"
)

// deflate input
func Encode(input string) (string, error) {
	result := ""
	lastChar := rune(0)
	count := 0
	for _, c := range input {
		if c >= '0' && c <= '9' {
			return "", errors.New("input must not contain numbers")
		} else if c == lastChar {
			count++
		} else {
			if count > 0 {
				result += strconv.Itoa(count) + string(lastChar)
			}
			lastChar = c
			count = 1
		}
	}
	if count > 0 {
		result += strconv.Itoa(count) + string(lastChar)
	}
	return result, nil
}
