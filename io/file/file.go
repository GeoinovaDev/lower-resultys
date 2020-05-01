package file

import (
	"io/ioutil"
	"os"
)

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

// Exist ...
func Exist(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
