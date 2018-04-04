package exec

import (
	"git.resultys.com.br/framework/lower/log"
)

func Trying(tentativas int, code func(), err func(string), finish func()) {

	for i := 0; i < tentativas; i++ {
		Try(func() {
			code()
			i = 10000
		}).Catch(func(message string) {
			log.Logger.Save(message, log.WARNING)
			err(message)
		})
	}

	finish()
}
