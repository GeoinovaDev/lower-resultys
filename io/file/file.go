package file

import "io/ioutil"

// GetContent ...
func GetContent(filename string) string {
	dat, err := ioutil.ReadFile(filename)

	if err != nil {
		return ""
	}

	return string(dat)
}

// WriteContent ...
func WriteContent(filename string, content string) {
	ioutil.WriteFile(filename, []byte(content), 0644)
}
