package exception

import (
	"runtime"

	"git.resultys.com.br/lib/lower/syslog"
	"git.resultys.com.br/lib/lower/time/datetime"
)

// Tipo ...
const (
	WARNING  = "WR"
	PANIC    = "PA"
	CRITICAL = "CR"
)

// Stack ...
type Stack struct {
	Filename string `json:"filename" bson:"filename"`
	Line     int    `json:"line" bson:"line"`
}

// Exception ...
type Exception struct {
	Message  string  `json:"message" bson:"message"`
	CreateAt string  `json:"create_at" bson:"create_at"`
	Tipo     string  `json:"tipo" bson:"tipo"`
	Stack    []Stack `json:"stack" bson:"stack"`
}

// New ...
func New(message string, tipo string) *Exception {
	return &Exception{
		Message:  message,
		CreateAt: datetime.Now().String(),
		Stack:    []Stack{},
		Tipo:     tipo,
	}
}

// AddStack ...
func (e *Exception) AddStack(filename string, line int) {
	e.Stack = append(e.Stack, Stack{
		Filename: filename,
		Line:     line,
	})
}

// Raise ...
func Raise(message string, tipo string) *Exception {
	ex := New(message, tipo)

	for i := 0; ; i++ {
		_, fn, ln, ok := runtime.Caller(i)
		if !ok {
			break
		}

		ex.AddStack(fn, ln)
	}

	syslog.Get().Save(ex)

	return ex
}
