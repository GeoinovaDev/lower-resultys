package net

import (
	"git.resultys.com.br/framework/lower/library/convert"
)

type Protocol struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func (p *Protocol) ToJson() string {
	return convert.JsonToString(p)
}
