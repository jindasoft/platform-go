package xutils

import "strings"

func GetSlug(str string) string {
	lower := strings.ToLower(str)
	word := strings.Fields(lower)

	return strings.Join(word, "-")
}
