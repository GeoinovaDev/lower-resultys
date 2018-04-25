package mongo

import (
	"time"

	"git.resultys.com.br/lib/lower/db/cstring"
	"git.resultys.com.br/lib/lower/exec"
	"git.resultys.com.br/lib/lower/log"
	"git.resultys.com.br/lib/lower/net/loopback"
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
	exec.Try(func() {
		mongo.createConn()
		if mongo.session == nil {
			return
		}
		defer mongo.session.Close()

		c := mongo.session.DB(mongo.db).C(mongo.c)
		query(c)
	}).Catch(func(message string) {
		log.Logger.Save(message, log.WARNING, loopback.IP())
	})

	return mongo
}

func (mongo *Mongo) createConn() *Mongo {
	conn := cstring.Get("mongo")
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{conn.Host},
		Timeout:  60 * time.Second,
		Database: conn.Db,
		Username: conn.User,
		Password: conn.Pass,
	}

	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		mongo.Error = err
		log.Logger.Save(err.Error(), log.WARNING, loopback.IP())
		mongo.session = nil
		return mongo
	}

	mongo.Error = nil
	mongo.session = session

	return mongo
}
