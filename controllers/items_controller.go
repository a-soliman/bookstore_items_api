package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/a-soliman/bookstore_items_api/domain/items"
	"github.com/a-soliman/bookstore_items_api/services"
	"github.com/a-soliman/bookstore_items_api/utils/http_utils"
	"github.com/a-soliman/bookstore_oauth-go/oauth"
	"github.com/a-soliman/bookstore_utils-go/rest_errors"
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
	// authorize the request
	if err := oauth.AuthenticateRequest(r); err != nil {
		http_utils.ResponseError(w, err)
		return
	}

	// read the req body
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respError := rest_errors.NewBadRequestError("invalid request body")
		http_utils.ResponseError(w, respError)
		return
	}
	defer r.Body.Close()

	// build the itemRequest json
	var itemRequest items.Item
	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		respError := rest_errors.NewBadRequestError("invalid json body")
		http_utils.ResponseError(w, respError)
		return
	}

	// append the seller id
	itemRequest.Seller = oauth.GetCallerID(r)

	// create the item
	result, createErr := services.ItemsService.Create(itemRequest)
	if err != nil {
		http_utils.ResponseError(w, createErr)
		return
	}
	http_utils.ResponseJSON(w, http.StatusCreated, result)
}
