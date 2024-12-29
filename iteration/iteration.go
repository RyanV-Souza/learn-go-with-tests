package iteration

import "strings"

func Repeat(character string, count int) string {
	if count <= 0 {
		count = 5
	}

	return strings.Repeat(character, count)
}
