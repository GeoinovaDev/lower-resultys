package pgsql

import (
	"database/sql"
	"log"

	"git.resultys.com.br/lib/lower/exception"
	"git.resultys.com.br/lib/lower/exec/try"
)

// PGSql struct
type PGSql struct {
	cs         string
	db         *sql.DB
	tentativas int
}

var current map[string]*PGSql

// New ...
func New(cs string) *PGSql {
	return &PGSql{
		cs:         cs,
		tentativas: 3,
	}
}

// SetTentativas ...
func (pg *PGSql) SetTentativas(tentativas int) {
	pg.tentativas = tentativas
}

// Connect ...
func (pg *PGSql) Connect() error {
	db, err := sql.Open("postgres", pg.cs)

	pg.db = db

	return err
}

// Close ...
func (pg *PGSql) Close() {
	if pg.db != nil {
		pg.db.Close()
	}
}

// IsAlive ...
func (pg *PGSql) IsAlive() bool {
	if pg.db == nil {
		return false
	}

	return pg.db.Ping() == nil
}

// GetInstance ...
func GetInstance(cs string) *PGSql {
	if current == nil {
		current = make(map[string]*PGSql)
	}

	if _, ok := current[cs]; !ok {
		current[cs] = New(cs)
	}

	if !current[cs].IsAlive() {
		current[cs].Connect()
	}

	return current[cs]
}

// Query ...
func (pg *PGSql) Query(query func(*sql.DB)) {
	try.New().SetTentativas(pg.tentativas).Run(func() {
		query(pg.db)
	}).Catch(func(m string) {
		log.Println(m)
		exception.Raise(m, exception.WARNING)
	})
}
