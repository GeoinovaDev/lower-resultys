package exec

import (
	"log"
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
			switch err.(type) {
			case string:
				trying.err = err.(string)
			case []string:
				trying.err = strings.Join(err.([]string), ". ")
			default:
				trying.err = "erro de runtime"
			}

			trying.throw = true
			t = trying

			log.Println(err)

			if trying.cacheCatch {
				trying.cbCatch(trying.err)
			}

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
