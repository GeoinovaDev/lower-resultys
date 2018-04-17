package exec

import (
	"time"

	"git.resultys.com.br/lib/lower/log"
	"git.resultys.com.br/lib/lower/net/loopback"
)

// Trying tenta executar a função em até 'tentativas' vezes.
// Caso ocorra successo chama o callback success
// Caso ocorra error chama o callback error
// Ao final do processo chama o callback finish
func Trying(tentativas int, code func(), success func(), err func(), finish func()) {
	b := false
	for i := 0; i < tentativas; i++ {
		Try(func() {
			code()
			i = 10000
			b = true
		}).Catch(func(message string) {
			log.Logger.Save(message, log.WARNING, loopback.IP())
			time.Sleep(5 * time.Second)
		})
	}

	if b {
		success()
	} else {
		err()
	}

	finish()
}
