package net

import (
	"git.resultys.com.br/lib/lower/convert"
)

// Protocol é o protocol utilizado para comunicação via rede contendo
// Code, Status, Data, Message
type Protocol struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// ToJSON converte o protocolo em string
func (p *Protocol) ToJSON() string {
	return convert.JSONToString(p)
}

// Success codifica protocolo como sucesso
func Success(data interface{}) string {
	protocol := &Protocol{
		Code:   200,
		Status: "ok",
		Data:   data,
	}

	return protocol.ToJSON()
}

// Error codifica protocolo como error
func Error(code int, message string) string {
	protocol := &Protocol{
		Code:    code,
		Status:  "error",
		Data:    nil,
		Message: message,
	}

	return protocol.ToJSON()
}
