package config

import (
	"encoding/json"
	"git.resultys.com.br/framework/lower/log"
	"git.resultys.com.br/framework/lower/net/loopback"
	"io/ioutil"
)

var File = "./config.json"

func readFile() []byte {
	raw, err := ioutil.ReadFile(File)
	if err != nil {
		log.Logger.Save("não foi possivel ler o arquivo config.json", log.PANIC, loopback.IP())
		raw = make([]byte, 0)
	}

	return raw
}

func Exist() bool {
	_, err := ioutil.ReadFile(File)
	if err != nil {
		return false
	}

	return true
}

func Save(obj interface{}) error {
	data, err := json.Marshal(obj)
	if err != nil {
		log.Logger.Save("não foi possivel salvar no arquivo config.json", log.WARNING, loopback.IP())
		return err
	}

	ioutil.WriteFile(File, data, 755)

	return nil
}

func Get(key string) string {
	var obj map[string]string

	raw := readFile()
	json.Unmarshal(raw, &obj)

	return obj[key]
}

func LoadInto(to interface{}) {
	raw := readFile()
	json.Unmarshal(raw, &to)
}
