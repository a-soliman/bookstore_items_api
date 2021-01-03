package app

import (
	"net/http"

	"github.com/a-soliman/bookstore_items_api/controllers"
)

func mapUrls() {
	router.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodGet)

	router.HandleFunc("/items/:item_id", controllers.ItemsController.Get).Methods(http.MethodGet)
	router.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)
}