package encode

import (
	"encoding/json"

	"git.resultys.com.br/lib/lower/exception"
)

// JSON encode
func JSON(obj interface{}) string {
	_json, err := json.Marshal(obj)
	if err != nil {
		exception.Raise(err.Error(), exception.WARNING)
		return ""
	}

	return string(_json)
}
