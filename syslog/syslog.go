package syslog

// ILogger ...
type ILogger interface {
	Save(interface{})
}

// DefaultLogger ...
type DefaultLogger struct {
}

// log ...
var log ILogger = DefaultLogger{}

// Set ...
func Set(logger ILogger) {
	log = logger
}

// Get ...
func Get() ILogger {
	return log
}

// Save ...
func (d DefaultLogger) Save(message interface{}) {

}
