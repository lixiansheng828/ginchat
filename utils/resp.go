package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H struct {
	Code  int
	Msg   string
	Data  interface{}
	Rows  interface{}
	Total interface{}
}

func Resp(w http.ResponseWriter, code int, data interface{}, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	h := H{
		Code:  code,
		Data:  data,
		Msg: msg,
	}
	ret, err := json.Marshal(h)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(ret)
}

func ResList(w http.ResponseWriter, code int, data interface{}, totla interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	h := H{
		Code:  code,
		Rows:  data,
		Total: totla,
	}
	ret, err := json.Marshal(h)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(ret)
}

func ResFail(w http.ResponseWriter, msg string) {
	Resp(w, -1, nil, msg)
}
func ResOK(w http.ResponseWriter, data interface{}, msg string) {
	Resp(w, 0, data, msg)
}

func ResOKList(w http.ResponseWriter, data interface{}, total interface{}) {
	ResList(w, 0, data, total)
}
