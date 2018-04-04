package db

import (
	"git.resultys.com.br/framework/lower/str"
)

type ConnectionString struct {
	Host string
	User string
	Pass string
	Db   string
}

func (c *ConnectionString) Get() string {
	return str.Format("host={0} user={1} password={2} dbname={3} sslmode=disable",
		c.Host,
		c.User,
		c.Pass,
		c.Db,
	)
}
