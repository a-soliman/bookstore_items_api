package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

var (
	router *mux.Router = mux.NewRouter()
)

// StartApplication starts the app
func StartApplication() {
	mapUrls()

	server := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
