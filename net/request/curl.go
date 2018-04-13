package request

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"git.resultys.com.br/lib/lower/log"
	"git.resultys.com.br/lib/lower/net/loopback"
)

// ProxyURL contem o endereco do proxy no formato: http://dominio
var ProxyURL = ""

// CURL estrutura da request
type CURL struct {
	url     string
	body    string
	request *http.Request
	headers map[string]string
}

// URLEncode retorna a string codificada em url
func URLEncode(str string) string {
	return url.PathEscape(str)
}

// URLDecode retorna a string decodifica em url
func URLDecode(str string) string {
	s, _ := url.PathUnescape(str)
	return s
}

// Create cria uma request
func Create(url string) *CURL {
	curl := &CURL{url: url}
	curl.headers = make(map[string]string)
	return curl
}

// AddHeader adiciona cabecalho ao HTTP
func (curl *CURL) AddHeader(key string, value string) *CURL {
	curl.headers[key] = value
	return curl
}

// Get faz um request GET
// Retorna o body como string e o error
// Salva no sistema de log caso ocorra erro
func (curl *CURL) Get() (string, error) {
	err := curl.createRequest("GET", "")
	if err != nil {
		return "", err
	}

	curl.injectHeaders()
	body, err := curl.sendRequest()
	if err != nil {
		return "", err
	}

	return body, nil
}

// GetJSON faz request GET para url informada.
// Injeta o retorno no parametro obj
// Retorna error ou nil
func (curl *CURL) GetJSON(obj interface{}) error {
	text, err := curl.Get()
	if err != nil {
		return err
	}

	json.Unmarshal([]byte(text), &obj)

	return nil
}

// Post faz uma requisição POST para url informada, submetendo o dados tipo form-urlencoded
// Retorna a resposta em string e error
func (curl *CURL) Post(formData map[string]string) (string, error) {
	values := url.Values{}
	for key, value := range formData {
		values.Add(key, value)
	}

	err := curl.createRequest("POST", values.Encode())
	if err != nil {
		return "", err
	}

	curl.request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	curl.injectHeaders()
	body, err := curl.sendRequest()
	if err != nil {
		return "", err
	}

	return body, nil
}

// PostJSON faz um requisição POST para url informada convertendo o objeto passado por parametro em json e com o cabeçalho do tipo application/json
// Retorna o body como string e o error
func (curl *CURL) PostJSON(obj interface{}) (string, error) {
	_json, err := json.Marshal(obj)
	if err != nil {
		log.Logger.Save(err.Error(), log.WARNING, loopback.IP())
		return "", errors.New("erro ao codificar o json")
	}

	err = curl.createRequest("POST", string(_json))
	if err != nil {
		return "", err
	}

	curl.request.Header.Set("Content-Type", "application/json")
	curl.injectHeaders()
	body, err := curl.sendRequest()
	if err != nil {
		return "", err
	}

	return body, nil
}

func (curl *CURL) createRequest(method string, data string) error {
	var req *http.Request
	var err error

	if len(data) == 0 {
		req, err = http.NewRequest(method, curl.url, nil)
	} else {
		req, err = http.NewRequest(method, curl.url, strings.NewReader(data))
	}

	if err != nil {
		log.Logger.Save(err.Error(), log.WARNING, loopback.IP())
		return errors.New("erro ao criar a request")
	}

	curl.request = req
	return nil
}

func (curl *CURL) sendRequest() (string, error) {
	resp, err := createClient().Do(curl.request)
	if err != nil {
		log.Logger.Save(err.Error(), log.WARNING, loopback.IP())
		return "", errors.New("error ao conectar a url")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Logger.Save(err.Error(), log.WARNING, loopback.IP())
		return "", errors.New("erro ao ler o conteudo do body")
	}

	curl.body = string(body)

	return string(body), nil
}

func (curl *CURL) injectHeaders() {
	for key, value := range curl.headers {
		curl.request.Header.Set(key, value)
	}
}

func createClient() *http.Client {
	var transport = &http.Transport{
		MaxIdleConns: 100,
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}

	if len(ProxyURL) > 5 {
		urlProxy, err := url.Parse(ProxyURL)
		if err != nil {
			log.Logger.Save(err.Error(), log.PANIC, loopback.IP())
		} else {
			transport.Proxy = http.ProxyURL(urlProxy)
		}
	}

	return &http.Client{
		Timeout:   10 * time.Second,
		Transport: transport,
	}
}
