package convert

import (
	"encoding/json"

	"git.resultys.com.br/framework/lower/log"
	"git.resultys.com.br/framework/lower/net/loopback"
)

// JSONToString converte um json em string
func JSONToString(obj interface{}) string {
	_json, err := json.Marshal(obj)
	if err != nil {
		log.Logger.Save(err.Error(), log.WARNING, loopback.IP())
		return ""
	}

	return string(_json)
}

// BytesToJSON converte slice de bytes em objeto
func BytesToJSON(bytes []byte, obj interface{}) {
	json.Unmarshal(bytes, &obj)
}

// StringToJSON converte string em objeto
func StringToJSON(str string, obj interface{}) {
	json.Unmarshal([]byte(str), &obj)
}
