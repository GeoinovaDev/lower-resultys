package mongo

import (
	"time"

	"git.resultys.com.br/lib/lower/db/cstring"
	"git.resultys.com.br/lib/lower/log"
	"git.resultys.com.br/lib/lower/net/loopback"
	mgo "gopkg.in/mgo.v2"
)

// New cria um sessão com o mongo
func New() *mgo.Session {
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
		log.Logger.Save(err.Error(), log.WARNING, loopback.IP())
		return nil
	}

	return session
}
