package server

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"git.resultys.com.br/lib/lower/config"
	"git.resultys.com.br/lib/lower/exception"
	"git.resultys.com.br/lib/lower/exec"
)

// Port contém informação sobre a porta utilizada pelo servidor
var Port = ":80"
var listing = false

// QueryString contém a estrutura dos valores passados por parametro na url
type QueryString struct {
	values url.Values
}

// Get Retorna um valor para chave na query string
func (qs QueryString) Get(key string) string {
	if val, ok := qs.values[key]; ok {
		return val[0]
	}

	return ""
}

// GetInt ...
func (qs QueryString) GetInt(key string) int {
	value := qs.Get(key)
	v, err := strconv.Atoi(value)

	if err != nil {
		return 0
	}

	return v
}

func createServer() *http.Server {
	port := Port

	if config.Exist() {
		port = config.Get("port")
	}

	if len(port) == 0 {
		port = Port
	}

	return &http.Server{
		Addr: port,
	}
}

// OnGet possui callback de rota para requisições do tipo GET
func OnGet(route string, handler func(QueryString) string) {
	http.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		qs := QueryString{r.URL.Query()}

		defer r.Body.Close()

		exec.Try(func() {

			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, x-requested-with")

			text := handler(qs)
			fmt.Fprint(w, text)
		})
	})
}

// OnPost possui callback de rota para requisições do tipo POST
func OnPost(route string, handler func(QueryString, string) string) {
	http.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		data := buf.String()
		qs := QueryString{r.URL.Query()}

		exec.Try(func() {

			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

			text := handler(qs, data)
			fmt.Fprint(w, text)
		})
	})
}

// On possui callback de rota para requisições de qualquer metodo
func On(route string, handler func() string) {
	http.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		exec.Try(func() {
			text := handler()
			fmt.Fprint(w, text)
		})
	})
}

// Start inicia o serviço
func Start() {
	if listing {
		exception.Raise("serviço já foi iniciado", exception.WARNING)
		return
	}

	server := createServer()
	err := server.ListenAndServe()
	if err != nil {
		println(err.Error())
		fmt.Scanln()
		panic(err)
	}
}
