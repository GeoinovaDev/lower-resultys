package server

import (
	"bytes"
	"fmt"
	"git.resultys.com.br/framework/lower/library/config"
	"git.resultys.com.br/framework/lower/log"
	"net/http"
	"net/url"
)

var Port = ":80"

var listing bool = false

type QueryString struct {
	values url.Values
}

func (qs QueryString) Get(key string) string {
	return qs.values[key][0]
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

func OnGet(route string, handler func(QueryString) string) {
	http.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		defer func() {
			err := recover()
			if err != nil {
				log.Logger.Save(fmt.Sprint(err), log.WARNING)
			}
		}()

		text := handler(QueryString{r.URL.Query()})
		fmt.Fprint(w, text)
	})
}

func OnPost(route string, handler func(QueryString, string) string) {
	http.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		defer func() {
			err := recover()
			if err != nil {
				log.Logger.Save(fmt.Sprint(err), log.WARNING)
			}
		}()

		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)

		text := handler(QueryString{r.URL.Query()}, buf.String())
		fmt.Fprint(w, text)
	})
}

func On(route string, handler func() string) {
	http.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		defer func() {
			err := recover()
			if err != nil {
				log.Logger.Save(fmt.Sprint(err), log.WARNING)
			}
		}()

		text := handler()
		fmt.Fprint(w, text)
	})
}

func Start() {
	if listing {
		log.Logger.Save("servidor ja esta em execucao", log.WARNING)
		return
	}

	server := createServer()
	server.ListenAndServe()
}
