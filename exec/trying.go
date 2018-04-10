package exec

import (
	"git.resultys.com.br/framework/lower/log"
	"git.resultys.com.br/framework/lower/net/loopback"
)

// Trying tenta executar a função em até 'tentativas' vezes.
// Caso ocorra erro chama o callback err
// Ao final da execução com erro ou não invoka o callback finish
func Trying(tentativas int, code func(), err func(string), finish func()) {

	for i := 0; i < tentativas; i++ {
		Try(func() {
			code()
			i = 10000
		}).Catch(func(message string) {
			log.Logger.Save(message, log.WARNING, loopback.IP())
			err(message)
		})
	}

	finish()
}
