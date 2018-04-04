package config

import (
	"encoding/json"
	"git.resultys.com.br/framework/lower/log"
	"io/ioutil"
)

var File = "./config.json"

func readFile() []byte {
	raw, err := ioutil.ReadFile(File)
	if err != nil {
		log.Logger.Save("não foi possivel ler o arquivo config.json", log.PANIC)
		panic("não foi possivel ler o arquivo config.json")
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
