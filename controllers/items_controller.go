package controllers

import (
	"fmt"
	"net/http"

	"github.com/a-soliman/bookstore_items_api/domain/items"
	"github.com/a-soliman/bookstore_items_api/services"
	"github.com/a-soliman/bookstore_oauth-go/oauth"
)

var (
	// ItemsController the exported instance
	ItemsController ItemsControllerInterface = &itemsController{}
)

// ItemsControllerInterface the items controller interface
type ItemsControllerInterface interface {
	Get(http.ResponseWriter, *http.Request)
	Create(http.ResponseWriter, *http.Request)
}

type itemsController struct{}

// Get gets an item by id if exists
func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {}

// Create creates an item
func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		// TODO: Return error to the caller
		return
	}
	item := items.Item{
		Seller: oauth.GetCallerID(r),
	}

	result, err := services.ItemsService.Create(item)
	if err != nil {
		// TODO: Return error json to the user.
	}
	fmt.Println(result)
	// Todo return the created item with

}
