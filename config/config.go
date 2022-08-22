package config

import (
	"encoding/json"
	"io/ioutil"
	"sync"

	"github.com/GeoinovaDev/lower-resultys/exception"
)

// File variável global contendo o filename default do config.json
var File = "./config.json"
var mutex = &sync.Mutex{}

func readFile() []byte {
	raw, err := ioutil.ReadFile(File)
	if err != nil {
		exception.Raise("não foi possivel ler o arquivo config.json", exception.WARNING)
		raw = make([]byte, 0)
	}

	return raw
}

// Exist verifica se o config.json existe
func Exist() bool {
	_, err := ioutil.ReadFile(File)
	if err != nil {
		return false
	}

	return true
}

// Save salva um objeto de configuração no config.json
func Save(obj interface{}) error {
	mutex.Lock()
	defer mutex.Unlock()

	data, err := json.Marshal(obj)
	if err != nil {
		exception.Raise("não foi possivel salvar no arquivo config.json", exception.WARNING)
		return err
	}

	ioutil.WriteFile(File, data, 755)
	return nil
}

// Get retorna uma propriedade do config.json
func Get(key string) string {
	var obj map[string]string

	raw := readFile()
	json.Unmarshal(raw, &obj)

	return obj[key]
}

// LoadInto carrega o config.json em um objeto
func LoadInto(to interface{}) {
	raw := readFile()
	json.Unmarshal(raw, &to)
}
