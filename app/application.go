package app

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var (
	router *mux.Router = mux.NewRouter()
)

// StartApplication starts the app
func StartApplication() {
	mapUrls()

	server := &http.Server{
		Handler:      router,
		Addr:         "localhost:8000",
		WriteTimeout: 500 * time.Millisecond,
		ReadTimeout:  2 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}