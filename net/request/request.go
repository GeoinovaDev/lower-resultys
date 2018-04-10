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

	"git.resultys.com.br/framework/lower/log"
	"git.resultys.com.br/framework/lower/net/loopback"
)

// ProxyURL contem o endereco do proxy no formato: http://dominio
var ProxyURL = ""

// URLEncode retorna a string codificada em url
func URLEncode(str string) string {
	return url.PathEscape(str)
}

// URLDecode retorna a string decodifica em url
func URLDecode(str string) string {
	s, _ := url.PathUnescape(str)
	return s
}

// Get faz um request GET para a url informada
// Retorna o body como string e o error
// Salva no sistema de log caso ocorra erro
func Get(url string) (text string, erro error) {
	resp, err1 := createClient().Get(url)
	if err1 != nil {
		log.Logger.Save(err1.Error(), log.WARNING, loopback.IP())
		return "", errors.New("error ao conectar a url")
	}
	defer resp.Body.Close()

	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		log.Logger.Save(err2.Error(), log.WARNING, loopback.IP())
		return "", errors.New("erro ao ler o conteudo do body")
	}

	return string(body), nil
}

// GetJSON faz request GET para url informada.
// Injeta o retorno no parametro obj
// Retorna error ou nil
func GetJSON(url string, obj interface{}) error {
	text, err := Get(url)
	if err != nil {
		return err
	}

	json.Unmarshal([]byte(text), &obj)

	return nil
}

// Post faz uma requisição POST para url informada, submetendo o dados tipo form-urlencoded
// Retorna a resposta em string e error
func Post(url string, formData url.Values) (string, error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(formData.Encode()))
	if err != nil {
		log.Logger.Save(err.Error(), log.WARNING, loopback.IP())
		return "", errors.New("erro ao criar o post")
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := createClient().Do(req)
	if err != nil {
		log.Logger.Save(err.Error(), log.WARNING, loopback.IP())
		return "", errors.New("erro ao conectar ao servidor")
	}

	body, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		log.Logger.Save(err1.Error(), log.WARNING, loopback.IP())
		return "", errors.New("erro ao ler os dados de retorno do post")
	}

	return string(body), nil
}

// PostJSON faz um requisição POST para url informada convertendo o objeto passado por parametro em json e com o cabeçalho do tipo application/json
// Retorna o body como string e o error
func PostJSON(url string, obj interface{}) (string, error) {
	_json, err := json.Marshal(obj)
	if err != nil {
		log.Logger.Save(err.Error(), log.WARNING, loopback.IP())
		return "", errors.New("erro ao codificar o json")
	}

	req, err := http.NewRequest("POST", url, strings.NewReader(string(_json)))
	if err != nil {
		log.Logger.Save(err.Error(), log.WARNING, loopback.IP())
		return "", errors.New("erro ao criar o post")
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := createClient().Do(req)
	if err != nil {
		log.Logger.Save(err.Error(), log.WARNING, loopback.IP())
		return "", errors.New("erro ao conectar ao servidor")
	}

	body, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		log.Logger.Save(err1.Error(), log.WARNING, loopback.IP())
		return "", errors.New("erro ao ler os dados de retorno do post")
	}

	return string(body), nil
}

func createClient() *http.Client {
	var transport = &http.Transport{
		MaxIdleConns: 100,
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}

	if len(ProxyURL) > 0 {
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
