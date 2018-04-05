package exec

import (
	"fmt"
	"git.resultys.com.br/framework/lower/log"
	"git.resultys.com.br/framework/lower/net/loopback"
	"strings"
)

type try struct {
	err   string
	throw bool

	cacheCatch bool
	cbCatch    func(string)
}

func Try(code func()) (t *try) {
	trying := &try{}

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

			log.Logger.Save(trying.err, log.WARNING, loopback.IP())

			return
		}
	}()

	code()

	return trying
}

func (t *try) Catch(code func(string)) {
	t.cbCatch = code
	t.cacheCatch = true

	if t.throw {
		code(t.err)
	}
}
