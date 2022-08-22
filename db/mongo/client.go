package mongo

import (
	"time"

	"github.com/GeoinovaDev/lower-resultys/db/cstring"
	"github.com/GeoinovaDev/lower-resultys/exception"
	"github.com/GeoinovaDev/lower-resultys/exec"
	mgo "gopkg.in/mgo.v2"
)

// Mongo struct
type Mongo struct {
	Error error

	db      string
	c       string
	session *mgo.Session
}

// New cria um sessão com o mongo
func New() *Mongo {
	return &Mongo{}
}

// DB seta o banco
func (mongo *Mongo) DB(db string) *Mongo {
	mongo.db = db

	return mongo
}

// C seta a colecao
func (mongo *Mongo) C(c string) *Mongo {
	mongo.c = c

	return mongo
}

// Query executa uma consulta
func (mongo *Mongo) Query(query func(*mgo.Collection)) *Mongo {
	exec.Tryx(5, func() {
		mongo.createConn()
		if mongo.session == nil {
			return
		}
		defer mongo.session.Close()

		mongo.session.SetSocketTimeout(12 * time.Second)

		c := mongo.session.DB(mongo.db).C(mongo.c)
		query(c)
	})

	return mongo
}

func (mongo *Mongo) createConn() *Mongo {
	conn := cstring.Get("mongo")
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{conn.Host},
		Timeout:  10 * time.Second,
		Database: conn.Db,
		Username: conn.User,
		Password: conn.Pass,
	}

	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		mongo.Error = err
		exception.Raise(err.Error(), exception.WARNING)
		mongo.session = nil
		return mongo
	}

	mongo.Error = nil
	mongo.session = session

	return mongo
}
