package logfile

import (
	"io/ioutil"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/GeoinovaDev/lower-resultys/str"
)

// Log ...
type Log struct {
	filename string
	mx       *sync.Mutex
	id       int
	limit    int
	count    int
}

var logfiles map[string]*Log

// New ...
func New(filename string) *Log {
	return &Log{
		mx:       &sync.Mutex{},
		filename: filename,
		limit:    0,
		count:    0,
		id:       1,
	}
}

// Limit ...
func (log *Log) Limit(limit int) *Log {
	log.limit = limit

	return log
}

// GetInstance ...
func GetInstance(name string) *Log {
	logfiles = _getInstance()

	if log, ok := logfiles[name]; ok {
		return log
	}

	logfiles[name] = New(name)

	return logfiles[name]
}

// Read ...
func (log *Log) Read() string {
	log.mx.Lock()
	defer log.mx.Unlock()

	return _read(log.filename)
}

// Add ...
func (log *Log) Add(content string) {
	log.mx.Lock()
	defer log.mx.Unlock()
	filename := log.filename

	log.count++
	if log.count > log.limit {
		log.count = 0
		log.id++
	}

	filename = log._formatFilename()

	current := _read(filename)
	ioutil.WriteFile(filename, []byte(current+"\r\n"+content), 0644)
}

func (log *Log) _formatFilename() string {
	file, ex := _extractFile(log.filename)
	d := time.Now()
	hoje := str.Format("{0}-{1}-{2}", strconv.Itoa(d.Day()), d.Month().String(), strconv.Itoa(d.Year()))
	return file + "." + strconv.Itoa(log.id) + "." + hoje + "." + ex
}

func _getInstance() map[string]*Log {
	if logfiles == nil {
		logfiles = map[string]*Log{}
	}

	return logfiles
}

func _read(filename string) string {
	dat, err := ioutil.ReadFile(filename)

	if err != nil {
		return ""
	}

	return string(dat)
}

func _extractFile(filename string) (string, string) {
	p := strings.Split(filename, ".")

	if len(p) != 2 {
		return filename, ""
	}

	return p[0], p[1]
}
