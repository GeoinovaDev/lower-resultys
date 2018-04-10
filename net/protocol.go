package net

import (
	"git.resultys.com.br/framework/lower/library/convert"
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
