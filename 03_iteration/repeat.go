package iteration

import "strings"

const repeatCount = 5

func Repeat(character string, count int) string {
	var repeated strings.Builder
	for range count {
		repeated.WriteString(character)
	}

	return repeated.String()
}
