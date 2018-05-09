package convert

import (
	"encoding/json"

	"git.resultys.com.br/lib/lower/convert/encode"
	"git.resultys.com.br/lib/lower/log"
	"git.resultys.com.br/lib/lower/net/loopback"
)

// HTTPBuildQuery ...
func HTTPBuildQuery(arr map[string]string) string {
	query := ""

	for k, v := range arr {
		query += k + "=" + v + "&"
	}

	return encode.URL(string(query[:len(query)-1]))
}

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

// ArrayInterfaceToArrayString convert array interface to array string
// Return array de string
func ArrayInterfaceToArrayString(arr []interface{}) []string {
	result := []string{}

	for i := 0; i < len(arr); i++ {
		result = append(result, arr[i].(string))
	}

	return result
}
