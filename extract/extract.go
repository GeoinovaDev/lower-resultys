package extract

import (
	"html"
	"regexp"
	"strings"
)

type extract struct {
	list []string
}

var regex map[string]*regexp.Regexp
var b = false

func init() {
	if b {
		return
	}

	regex = make(map[string]*regexp.Regexp)
	b = true
}

// New cria a estrutura de recorte
func New(content string) *extract {
	return &extract{list: []string{content}}
}

// First retorna o primeiro item recortado
func (e *extract) First() string {
	if len(e.list) == 0 {
		return ""
	}

	return e.list[0]
}

// ToArray converte para array
func (e *extract) ToArray() []string {
	return e.list
}

func compile(pattern string) *regexp.Regexp {
	if val, ok := regex[pattern]; ok {
		return val
	} else {
		c, _ := regexp.Compile(pattern)
		regex[pattern] = c

		return c
	}
}

// Regex executa uma ER sobre as list clipada
func (e *extract) Regex(pattern string) *extract {
	list := e.list
	e.list = []string{}

	for i := 0; i < len(list); i++ {
		c := compile(pattern)
		arr := c.FindAllString(list[i], -1)
		for j := 0; j < len(arr); j++ {
			e.list = append(e.list, arr[i])
		}
	}

	return e
}

// Clips recorta fragmentos dentro de conteudo
func (e *extract) Clips(parts ...string) *extract {
	contents := e.list
	e.list = []string{}
	if len(parts) < 2 {
		panic("str.IndexOf: numero de parametros incorreto")
	}

	for j := 0; j < len(contents); j++ {
		index := 0
		content := contents[j]
	loop:
		for {
			if index >= len(content) {
				break
			}

			for i := 0; i < len(parts)-1; i++ {
				part := parts[i]
				_index := strings.Index(string(content[index:]), part)
				if _index == -1 {
					break loop
				}
				index += _index + len(part)
			}

			if index == 0 {
				break
			}

			f := strings.Index(string(content[index:]), parts[len(parts)-1])
			if f == -1 {
				break
			}
			f += index

			str := string(content[index:f])
			str = html.UnescapeString(str)
			str = strings.Trim(str, " ")

			e.list = append(e.list, str)
			index = f
		}
	}

	return e
}
