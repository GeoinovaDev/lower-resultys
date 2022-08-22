package encode

import (
	"encoding/json"

	"github.com/GeoinovaDev/lower-resultys/exception"
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
