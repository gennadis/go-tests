package iteration

import "strings"

func Repeat(character string, count int) string {
	var repeated string

	for i := 0; i < count; i++ {
		repeated += character
	}

	return repeated
}

func Counter(word, character string) int {
	return strings.Count(word, character)
}
