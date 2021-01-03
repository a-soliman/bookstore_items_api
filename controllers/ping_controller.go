package controllers

import "net/http"

const (
	pong = "pong"
)

var (
	// PingController the exported instance
	PingController PingControllerInterface = &pingController{}
)

// PingControllerInterface the ping controller interface
type PingControllerInterface interface {
	Ping(http.ResponseWriter, *http.Request)
}

type pingController struct{}

func (c *pingController) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(pong))
}
