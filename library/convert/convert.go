package convert

import (
	"encoding/json"
)

func JsonToString(obj interface{}) string {
	_json, err := json.Marshal(obj)
	if err != nil {
		return ""
	}

	return string(_json)
}

func BytesToJson(bytes []byte, obj interface{}) {
	json.Unmarshal(bytes, &obj)
}

func StringToJson(str string, obj interface{}) {
	json.Unmarshal([]byte(str), &obj)
}
