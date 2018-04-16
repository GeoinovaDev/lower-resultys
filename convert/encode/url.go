package encode

import "net/url"

// URL retorna a string codificada em url
func URL(str string) string {
	return url.PathEscape(str)
}
