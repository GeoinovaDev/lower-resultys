package str

import (
	"strings"
)

// CountOccurrences ...
func CountOccurrences(str string, keyword string) int {
	count := 0
	i := 0

	if len(str) == 0 || len(keyword) == 0 {
		return 0
	}

	for {
		str = string(str[i:])
		i = strings.Index(str, keyword)
		if i > -1 {
			count++
			i++
		} else {
			break
		}
	}

	return count
}
