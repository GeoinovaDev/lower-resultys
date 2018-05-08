package decode

import (
	"bytes"
	"encoding/json"

	"git.resultys.com.br/lib/lower/log"
	"git.resultys.com.br/lib/lower/net/loopback"
)

// JSON decode into object
func JSON(str string, obj interface{}) {
	b := bytes.NewBufferString(str)
	err := json.NewDecoder(b).Decode(&obj)
	if err != nil {
		log.Logger.Save(err.Error(), log.WARNING, loopback.IP())
	}
}
