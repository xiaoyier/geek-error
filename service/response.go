package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type response struct {
	code    ErrCode
	message string
}

func ResponseSuccess(o http.ResponseWriter) {
	fmt.Fprintf(o, getResponseString(ErrSuccess))
}

func ResponseError(o http.ResponseWriter, code ErrCode) {
	fmt.Fprintf(o, getResponseString(code))
}

func getResponseString(code ErrCode) string {
	rsp := response{
		code:    code,
		message: code.Message(),
	}
	data, _ := json.Marshal(rsp)
	return string(data)
}
