package request

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"git.resultys.com.br/lib/lower/exception"
)

// ProxyURL contem o endereco do proxy no formato: http://dominio
var ProxyURL = ""

// CURL estrutura da request
type CURL struct {
	url            string
	body           string
	request        *http.Request
	headers        map[string]string
	timeout        time.Duration
	proxy          string
	Status         int
	raiseException bool
}

// New cria uma request
func New(url string) *CURL {
	curl := &CURL{url: url}
	curl.timeout = 60
	curl.raiseException = false
	curl.headers = make(map[string]string)
	return curl
}

// EnableRaiseException ...
func (curl *CURL) EnableRaiseException(enable bool) *CURL {
	curl.raiseException = enable
	return curl
}

// SetProxy ...
func (curl *CURL) SetProxy(proxy string) *CURL {
	curl.proxy = proxy

	return curl
}

// SetTimeout define timeout da resposta em segundos
func (curl *CURL) SetTimeout(timeout int) *CURL {
	curl.timeout = time.Duration(timeout)
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
		if curl.raiseException {
			exception.Raise(err.Error(), exception.WARNING)
		}

		return "", errors.New("erro ao codificar o json = " + err.Error())
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

// PostRaw ...
func (curl *CURL) PostRaw(data string) (string, error) {
	err := curl.createRequest("POST", data)
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

func (curl *CURL) createRequest(method string, data string) error {
	var req *http.Request
	var err error

	if len(data) == 0 {
		req, err = http.NewRequest(method, curl.url, nil)
	} else {
		req, err = http.NewRequest(method, curl.url, strings.NewReader(data))
	}

	if err != nil {
		if curl.raiseException {
			exception.Raise(err.Error(), exception.WARNING)
		}

		return errors.New("erro ao criar a request = " + err.Error())
	}

	curl.request = req
	curl.request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.181 Safari/537.36")
	curl.request.Header.Set("Connection", "close")
	curl.request.Close = true

	return nil
}

func (curl *CURL) sendRequest() (string, error) {
	resp, err := curl.createClient().Do(curl.request)
	if err != nil {
		if curl.raiseException {
			exception.Raise(err.Error(), exception.WARNING)
		}

		return "", errors.New("error ao conectar a url " + curl.url + " = " + err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		if curl.raiseException {
			exception.Raise(err.Error(), exception.WARNING)
		}

		return "", errors.New("erro ao ler o conteudo do body " + curl.url + " = " + err.Error())
	}

	curl.Status = resp.StatusCode

	if resp.StatusCode != 200 {
		// exception.Raise(err.Error(), exception.WARNING)
		return "", errors.New("error codigo " + strconv.Itoa(resp.StatusCode) + " ao conectar a url " + curl.url)
	}

	curl.body = string(body)

	return string(body), nil
}

func (curl *CURL) injectHeaders() {
	for key, value := range curl.headers {
		curl.request.Header.Set(key, value)
	}
}

func (curl *CURL) createClient() *http.Client {
	var transport = &http.Transport{
		MaxIdleConns: 1000,
		Dial: (&net.Dialer{
			Timeout: curl.timeout * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}

	_proxyURL := ""

	if len(ProxyURL) > 5 {
		_proxyURL = ProxyURL
	}

	if len(curl.proxy) > 5 {
		_proxyURL = curl.proxy
	}

	if len(_proxyURL) > 5 {
		urlProxy, err := url.Parse(_proxyURL)
		if err != nil {
			if curl.raiseException {
				exception.Raise(err.Error(), exception.WARNING)
			}

		} else {
			transport.Proxy = http.ProxyURL(urlProxy)
		}
	}

	return &http.Client{
		Timeout:   curl.timeout * time.Second,
		Transport: transport,
	}
}
