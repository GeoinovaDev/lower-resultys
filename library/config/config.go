package config

import (
	"encoding/json"
	"io/ioutil"
)

var File = "./config.json"

func readFile() []byte {
	raw, err := ioutil.ReadFile(File)
	if err != nil {
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
