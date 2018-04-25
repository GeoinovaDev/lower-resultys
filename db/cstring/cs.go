package cstring

import (
	"git.resultys.com.br/lib/lower/db"
)

var list map[string]*db.ConnectionString

func init() {
	list = make(map[string]*db.ConnectionString)
}

// Add adiciona connectionString
func Add(name string, cnn *db.ConnectionString) {
	list[name] = cnn
}

// Get retorna connection string
func Get(name string) *db.ConnectionString {
	return list[name]
}
