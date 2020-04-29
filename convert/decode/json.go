package decode

import (
	"bytes"
	"encoding/json"

	"git.resultys.com.br/lib/lower/exception"
)

// JSON decode into object
func JSON(str string, obj interface{}) (string, bool) {
	b := bytes.NewBufferString(str)
	err := json.NewDecoder(b).Decode(&obj)

	if err != nil {
		exception.Raise(err.Error(), exception.WARNING)
		return err.Error(), false
	}

	return "", true
}
