package try

import (
	"fmt"
	"strings"
	"time"

	"git.resultys.com.br/lib/lower/backoff"
)

// Tryx ...
type Tryx struct {
	err   string
	throw bool

	timeout    time.Duration
	backoff    bool
	tentativas int

	isSuccess bool

	cacheCatch bool
	cbCatch    func(string)
}

// New ...
func New() *Tryx {
	return &Tryx{backoff: false, tentativas: 1}
}

// SetTentativas ...
func (tx *Tryx) SetTentativas(tentativas int) *Tryx {
	tx.tentativas = tentativas
	return tx
}

// SetTimeout ...
func (tx *Tryx) SetTimeout(timeout time.Duration) *Tryx {
	tx.timeout = timeout
	return tx
}

// IsSuccess ...
func (tx *Tryx) IsSuccess() bool {
	return tx.isSuccess
}

// IsThrowException ...
func (tx *Tryx) IsThrowException() bool {
	return tx.throw
}

// ErrorMessage ...
func (tx *Tryx) ErrorMessage() string {
	return tx.err
}

// Run ...
func (tx *Tryx) Run(code func()) (t *Tryx) {
	b := &backoff.Backoff{
		Max: tx.timeout * time.Minute,
	}

	err := ""
	tx.throw = false

	for i := 0; i < tx.tentativas; i++ {
		tx.run(code, func() {
			i = tx.tentativas + 1
			err = ""
			tx.isSuccess = true
		}, func(message string) {
			err = message
			tx.isSuccess = false
			if tx.timeout > 0 {
				time.Sleep(b.Duration())
			}
		})
	}

	b.Reset()

	if len(err) > 0 {
		tx.err = err
		tx.throw = true

		if tx.cacheCatch {
			tx.cbCatch(tx.err)
		}
	}

	return tx
}

// Tryx ...
func (tx *Tryx) run(code func(), cbSuccess func(), cbErr func(string)) {
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
func (tx *Tryx) Catch(code func(string)) *Tryx {
	tx.cbCatch = code
	tx.cacheCatch = true

	if tx.throw {
		code(tx.err)
	}

	return tx
}
