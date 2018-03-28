package config

import (
	"encoding/json"
	"io/ioutil"
)

func readFile() []byte {
	raw, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic("não foi possivel ler o arquivo config.json")
	}

	return raw
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
