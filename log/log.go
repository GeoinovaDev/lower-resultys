package log

const (
	WARNING = 1
	PANIC   = 2
)

type ILogger interface {
	Save(string, int)
}

type DefaultLogger struct {
}

var Logger ILogger = DefaultLogger

func (d DefaultLogger) Save(message string, tpe int) {

}
