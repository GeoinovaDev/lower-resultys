package db

import (
	"github.com/GeoinovaDev/lower-resultys/str"
)

// ConnectionString é a estrutura contendo informações de acesso ao banco
type ConnectionString struct {
	Host string
	User string
	Pass string
	Db   string
}

// Get retorna a connection string
func (c *ConnectionString) Get() string {
	return str.Format("host={0} user={1} password={2} dbname={3} sslmode=disable",
		c.Host,
		c.User,
		c.Pass,
		c.Db,
	)
}
