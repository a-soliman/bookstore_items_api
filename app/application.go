package app

import (
	"net/http"
	"time"

	"github.com/a-soliman/bookstore_items_api/clients/elasticsearch"
	"github.com/gorilla/mux"
)

var (
	router *mux.Router = mux.NewRouter()
)

// StartApplication starts the app
func StartApplication() {
	elasticsearch.Init()

	mapUrls()

	server := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
