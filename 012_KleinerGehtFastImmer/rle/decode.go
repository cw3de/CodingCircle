package rle

import "errors"

// inflate input
func Decode(deflatedInput string) (string, error) {
	count := 0
	result := ""
	for _, c := range deflatedInput {
		if c >= '0' && c <= '9' {
			count = count*10 + int(c-'0')
		} else {
			if count == 0 {
				return "", errors.New("invalid input")
			}
			for i := 0; i < count; i++ {
				result += string(c)
			}
			count = 0
		}
	}
	return result, nil
}
