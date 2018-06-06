package exec

import (
	"fmt"
	"strings"
	"time"

	"git.resultys.com.br/lib/lower/backoff"
)

// TryxExec ...
type TryxExec struct {
	err   string
	throw bool

	cacheCatch bool
	cbCatch    func(string)
}

// Tryx ...
func Tryx(tentativas int, code func()) (t *TryxExec) {
	b := &backoff.Backoff{
		Max: 5 * time.Minute,
	}

	tryx := &TryxExec{}
	err := ""

	for i := 0; i < tentativas; i++ {
		tryx.run(code, func() {
			i = tentativas + 1
		}, func(message string) {
			err = message
			time.Sleep(b.Duration())
		})
	}

	b.Reset()

	if len(err) > 0 {
		tryx.err = err
		tryx.throw = true

		if tryx.cacheCatch {
			tryx.cbCatch(tryx.err)
		}
	}

	return tryx
}

func (tryx *TryxExec) run(code func(), cbSuccess func(), cbErr func(string)) {
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

			cbErr(msg)
		}
	}()

	code()

	cbSuccess()
}

// Catch ...
func (tryx *TryxExec) Catch(code func(string)) {
	tryx.cbCatch = code
	tryx.cacheCatch = true

	if tryx.throw {
		code(tryx.err)
	}
}
