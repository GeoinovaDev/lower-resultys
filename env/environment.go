package env

import (
	"git.resultys.com.br/lib/lower/io/file"
)

// Environment ...
type Environment struct {
	index   int
	current string
	items   []string
	path    string
}

var current *Environment

// New ...
func New() *Environment {
	return &Environment{
		items: []string{},
	}
}

// GetInstance ...
func GetInstance() *Environment {
	if current == nil {
		current = New()

		current.Add("production")
		current.Add("develop")
		current.Add("test")

		current.loadEnv()
	}

	return current
}

// Add ...
func (e *Environment) Add(name string) *Environment {
	e.items = append(e.items, name)

	return e
}

// Path ...
func (e *Environment) Path(path string) *Environment {
	e.path = path

	e.loadEnv()

	return e
}

// Set ...
func (e *Environment) Set(name string) *Environment {
	e.current = name
	e.index = e.getIndex(name)

	return e
}

// GetVar ...
func (e *Environment) GetVar(params ...interface{}) (interface{}, bool) {
	if e.isValid() && e.index < len(params) {
		return params[e.index], true
	}

	return nil, false
}

// GetVarString ...
func (e *Environment) GetVarString(params ...interface{}) string {
	v, b := e.GetVar(params...)
	if !b {
		return ""
	}

	return v.(string)
}

// GetVarInt ...
func (e *Environment) GetVarInt(params ...interface{}) (int, bool) {
	v, b := e.GetVar(params...)

	if !b {
		return -1, false
	}

	return v.(int), true
}

// Run ...
func (e *Environment) Run(params ...func()) bool {
	if e.isValid() && e.index < len(params) {
		params[e.index]()
		return true
	}

	return false
}

func (e *Environment) loadEnv() {
	if file.Exist(e.path + "/.__develop__") {
		e.Set("develop")
	} else if file.Exist(e.path + "/.__test__") {
		e.Set("test")
	} else {
		e.Set("production")
	}
}

func (e *Environment) isValid() bool {
	return e.index > -1
}

func (e *Environment) getIndex(name string) int {
	for i := 0; i < len(e.items); i++ {
		if e.items[i] == name {
			return i
		}
	}

	return -1
}
