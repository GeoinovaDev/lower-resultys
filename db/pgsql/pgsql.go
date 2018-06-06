package pgsql

import (
	"database/sql"
	"log"

	"git.resultys.com.br/lib/lower/exception"
	"git.resultys.com.br/lib/lower/exec"
)

// PGSql struct
type PGSql struct {
	cs string
}

// New ...
func New(cs string) *PGSql {
	return &PGSql{cs: cs}
}

// Query ...
func (pg *PGSql) Query(query func(*sql.DB)) {
	exec.Tryx(5, func() {
		db, err := sql.Open("postgres", pg.cs)
		if err != nil {
			panic(err)
		}
		defer db.Close()

		query(db)
	}).Catch(func(m string) {
		log.Println(m)
		exception.Raise(m, exception.WARNING)
	})
}
