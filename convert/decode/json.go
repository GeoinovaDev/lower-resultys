package decode

import (
	"bytes"
	"encoding/json"
)

// JSON decode into object
func JSON(str string, obj interface{}) {
	b := bytes.NewBufferString(str)
	json.NewDecoder(b).Decode(&obj)
}
