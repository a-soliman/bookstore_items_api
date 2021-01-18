package app

import (
	"log"
	"net/http"

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
	log.Fatal(http.ListenAndServe(":8000", router))
}
