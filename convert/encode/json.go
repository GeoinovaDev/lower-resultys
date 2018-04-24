package encode

import (
	"encoding/json"

	"git.resultys.com.br/lib/lower/log"
	"git.resultys.com.br/lib/lower/net/loopback"
)

// JSON encode
func JSON(obj interface{}) string {
	_json, err := json.Marshal(obj)
	if err != nil {
		log.Logger.Save(err.Error(), log.WARNING, loopback.IP())
		return ""
	}

	return string(_json)
}
