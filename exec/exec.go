package exec

import (
	"fmt"

	"git.resultys.com.br/framework/lower/log"
	"git.resultys.com.br/framework/lower/net/loopback"
)

// Loop executa infinitamente a função passada por parametro
func Loop(code func()) {
	While(func() bool {
		code()
		return true
	})
}

// While executa a função passada por parametro enquanto o valor de retorno for verdadeiro
func While(code func() bool) {
	for {
		ok := while(code)
		if ok == false {
			break
		}
	}
}

func while(code func() bool) (b bool) {
	defer func() {
		err := recover()
		if err != nil {
			log.Logger.Save(fmt.Sprint(err), log.WARNING, loopback.IP())
			b = true
			return
		}
	}()

	return code()
}
