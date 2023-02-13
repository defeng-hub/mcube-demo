package util

import (
	"encoding/json"
	mcubeException "github.com/infraboard/mcube/exception"
)

func DefaultException(code int, message string, data interface{}) *exception {
	return &exception{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
func NewException(code int, message string, data interface{}) *exception {
	return &exception{
		Namespace: "",
		HttpCode:  0,
		Code:      0,
		Reason:    "",
		Message:   "",
		Meta:      nil,
		Data:      nil,
	}
}

// APIException is impliment for api exception
type exception struct {
	Namespace string      `json:"namespace"`
	HttpCode  int         `json:"http_code"`
	Code      int         `json:"error_code"`
	Reason    string      `json:"reason"`
	Message   string      `json:"message"`
	Meta      interface{} `json:"meta"`
	Data      interface{} `json:"data"`
}

func (e *exception) ToJson() string {
	dj, _ := json.Marshal(e)
	return string(dj)
}

func (e *exception) Error() string {
	return e.Message
}

// ErrorCode Code exception's code, 如果code不存在返回-1
func (e *exception) ErrorCode() int {
	return int(e.Code)
}

func (e *exception) WithHttpCode(httpCode int) {
	e.HttpCode = httpCode
}

// GetHttpCode Code exception's code, 如果code不存在返回-1
func (e *exception) GetHttpCode() int {
	return int(e.HttpCode)
}

// WithMeta 携带一些额外信息
func (e *exception) WithMeta(m interface{}) mcubeException.APIException {
	e.Meta = m
	return e
}

func (e *exception) GetMeta() interface{} {
	return e.Meta
}

func (e *exception) WithData(d interface{}) mcubeException.APIException {
	e.Data = d
	return e
}

func (e *exception) GetData() interface{} {
	return e.Data
}

func (e *exception) Is(t error) bool {
	if v, ok := t.(mcubeException.APIException); ok {
		return e.ErrorCode() == v.ErrorCode()
	}

	return e.Message == t.Error()
}

func (e *exception) GetNamespace() string {
	return e.Namespace
}

func (e *exception) GetReason() string {
	return e.Reason
}

func (e *exception) WithNamespace(ns string) {
	e.Namespace = ns
}
