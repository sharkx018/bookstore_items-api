package controllers

import (
	"github.com/sharkx018/bookstore_items-api/utils/http_utils"
	"net/http"
)

var (
	PingController pingControllerInterface = &pingController{}
)

const (
	pong = "pong"
)

type pingControllerInterface interface {
	Ping(w http.ResponseWriter, r *http.Request)
}

type pingController struct{}

func (c *pingController) Ping(w http.ResponseWriter, r *http.Request) {
	person := struct {
		Name string `json:"name"`
		age  int    `json:"age"`
	}{
		Name: "Mukul",
		age:  25,
	}
	http_utils.RespondJson(w, http.StatusOK, person)
	//w.WriteHeader(http.StatusOK)
	//w.Write([]byte(pong))
}
