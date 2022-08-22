package exec

import (
	"fmt"
	"strings"

	"github.com/GeoinovaDev/lower-resultys/exception"
)

// TryExec é a estrutura contendo informações sobre a execução de uma função
type TryExec struct {
	err   string
	throw bool

	cacheCatch bool
	cbCatch    func(string)
}

// Try tenta executar a função passada por parametro não lançando excessão
func Try(code func()) (t *TryExec) {
	trying := &TryExec{}

	defer func() {
		err := recover()
		if err != nil {
			msg := ""
			switch err.(type) {
			case string:
				msg = err.(string)
			case []string:
				msg = strings.Join(err.([]string), ". ")
			case error:
				msg = fmt.Sprint(err)
			default:
				msg = "erro de runtime"
			}

			trying.err = msg
			trying.throw = true
			t = trying

			if trying.cacheCatch {
				trying.cbCatch(trying.err)
			}

			exception.Raise(msg, exception.WARNING)

			return
		}
	}()

	code()

	return trying
}

// Catch executa a função passada por parametro caso ocorreu um erro no Try
func (t *TryExec) Catch(code func(string)) {
	t.cbCatch = code
	t.cacheCatch = true

	if t.throw {
		code(t.err)
	}
}
