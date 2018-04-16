package decode

import (
	"net/url"
	"strings"
)

// URL retorna a string decodifica em url
func URL(str string) string {
	s, _ := url.PathUnescape(str)
	s = strings.Replace(s, "+", " ", -1)
	return s
}
