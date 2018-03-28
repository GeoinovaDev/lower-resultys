package request

import (
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type requestError struct {
	When time.Time
	What string
}

func (e requestError) Error() string {
	return e.What
}

var ProxyUrl string = ""

func createClient() *http.Client {
	var transport = &http.Transport{
		MaxIdleConns: 100,
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}

	if len(ProxyUrl) > 0 {
		urlProxy, err := url.Parse(ProxyUrl)
		if err != nil {
			panic("endereco do proxy esta errado")
		}

		transport.Proxy = http.ProxyURL(urlProxy)
	}

	return &http.Client{
		Timeout:   10 * time.Second,
		Transport: transport,
	}
}

func Get(url string) (text string, erro *requestError) {
	resp, err1 := createClient().Get(url)
	if err1 != nil {
		return "", &requestError{What: "error ao conectar a url"}
	}

	defer resp.Body.Close()
	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		return "", &requestError{What: "erro ao ler o conteudo do body"}
	}

	return string(body), nil
}

func GetJson(url string, obj interface{}) *requestError {
	text, err := Get(url)
	if err != nil {
		return err
	}

	json.Unmarshal([]byte(text), &obj)

	return nil
}

func Post(url string, formData url.Values) (string, *requestError) {
	req, err := http.NewRequest("POST", url, strings.NewReader(formData.Encode()))
	if err != nil {
		return "", &requestError{What: "erro ao criar o post"}
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := createClient().Do(req)
	if err != nil {
		return "", &requestError{What: "erro ao conectar ao servidor"}
	}

	body, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		return "", &requestError{What: "erro ao ler os dados de retorno do post"}
	}

	return string(body), nil
}

func PostJson(url string, obj interface{}) (string, *requestError) {
	_json, err := json.Marshal(obj)
	if err != nil {
		return "", &requestError{What: "erro ao codificar o json"}
	}

	req, err := http.NewRequest("POST", url, strings.NewReader(string(_json)))
	if err != nil {
		return "", &requestError{What: "erro ao criar o post"}
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := createClient().Do(req)
	if err != nil {
		return "", &requestError{What: "erro ao conectar ao servidor"}
	}

	body, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		return "", &requestError{What: "erro ao ler os dados de retorno do post"}
	}

	return string(body), nil
}
