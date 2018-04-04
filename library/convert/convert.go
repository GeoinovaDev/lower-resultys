package convert

import (
	"encoding/json"
	"git.resultys.com.br/framework/lower/log"
)

func JsonToString(obj interface{}) string {
	_json, err := json.Marshal(obj)
	if err != nil {
		log.Logger.Save(err.Error(), log.WARNING)
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
