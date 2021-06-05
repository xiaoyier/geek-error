package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type response struct {
	Code    ErrCode `json:"code"`
	Message string  `json:"message"`
}

func ResponseSuccess(o http.ResponseWriter) {
	fmt.Fprintf(o, getResponseString(ErrSuccess))
}

func ResponseError(o http.ResponseWriter, code ErrCode) {
	//fmt.Fprintf(o, getResponseString(code))
	fmt.Fprintf(o, getResponseString(code))
}

func getResponseString(code ErrCode) string {
	rsp := response{
		Code:    code,
		Message: code.Message(),
	}
	data, _ := json.Marshal(rsp)
	return string(data)
}
