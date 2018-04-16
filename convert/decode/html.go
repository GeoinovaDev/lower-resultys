package decode

import "html"

// HTML decodifica html entities
func HTML(str string) string {
	return html.UnescapeString(str)
}
