package server

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"

	"git.resultys.com.br/lib/lower/config"
	"git.resultys.com.br/lib/lower/log"
	"git.resultys.com.br/lib/lower/net/loopback"
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
		defer r.Body.Close()
		defer func() {
			err := recover()
			if err != nil {
				log.Logger.Save(fmt.Sprint(err), log.WARNING, loopback.IP())
			}
		}()

		text := handler(QueryString{r.URL.Query()})
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		fmt.Fprint(w, text)
	})
}

// OnPost possui callback de rota para requisições do tipo POST
func OnPost(route string, handler func(QueryString, string) string) {
	http.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		defer func() {
			err := recover()
			if err != nil {
				log.Logger.Save(fmt.Sprint(err), log.WARNING, loopback.IP())
			}
		}()

		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)

		text := handler(QueryString{r.URL.Query()}, buf.String())
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		fmt.Fprint(w, text)
	})
}

// On possui callback de rota para requisições de qualquer metodo
func On(route string, handler func() string) {
	http.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		defer func() {
			err := recover()
			if err != nil {
				log.Logger.Save(fmt.Sprint(err), log.WARNING, loopback.IP())
			}
		}()

		text := handler()
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		fmt.Fprint(w, text)
	})
}

// Start inicia o serviço
func Start() {
	if listing {
		log.Logger.Save("servidor ja esta em execucao", log.WARNING, loopback.IP())
		return
	}

	server := createServer()
	server.ListenAndServe()
}
