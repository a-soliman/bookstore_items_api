package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/a-soliman/bookstore_items_api/domain/items"
	"github.com/a-soliman/bookstore_items_api/domain/queries"
	"github.com/a-soliman/bookstore_items_api/services"
	"github.com/a-soliman/bookstore_items_api/utils/http_utils"
	"github.com/a-soliman/bookstore_oauth-go/oauth"
	"github.com/a-soliman/bookstore_utils-go/rest_errors"
	"github.com/gorilla/mux"
)

var (
	// ItemsController the exported instance
	ItemsController ItemsControllerInterface = &itemsController{}
)

// ItemsControllerInterface the items controller interface
type ItemsControllerInterface interface {
	Get(http.ResponseWriter, *http.Request)
	Create(http.ResponseWriter, *http.Request)
	Search(http.ResponseWriter, *http.Request)
}

type itemsController struct{}

// Get gets an item by id if exists
func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemID := strings.TrimSpace(vars["id"])

	item, err := services.ItemsService.Get(itemID)
	if err != nil {
		http_utils.ResponseError(w, err)
		return
	}

	http_utils.ResponseJSON(w, http.StatusOK, item)
}

// Create creates an item
func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	// authorize the request
	if err := oauth.AuthenticateRequest(r); err != nil {
		http_utils.ResponseError(w, err)
		return
	}

	sellerID := oauth.GetCallerID(r)

	if sellerID == 0 {
		unAuthorizedErr := rest_errors.NewUnauthorizedError("unable to retrieve user information from given access_token")
		http_utils.ResponseError(w, unAuthorizedErr)
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
	itemRequest.Seller = sellerID

	// create the item
	result, createErr := services.ItemsService.Create(itemRequest)
	if createErr != nil {
		http_utils.ResponseError(w, createErr)
		return
	}
	http_utils.ResponseJSON(w, http.StatusCreated, result)
}

func (c *itemsController) Search(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		apiErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.ResponseError(w, apiErr)
		return
	}
	defer r.Body.Close()

	var query queries.EsQuery
	if err := json.Unmarshal(bytes, &query); err != nil {
		apiErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.ResponseError(w, apiErr)
		return
	}

	items, searchErr := services.ItemsService.Search(query)
	if searchErr != nil {
		http_utils.ResponseError(w, searchErr)
		return
	}

	http_utils.ResponseJSON(w, http.StatusOK, items)
}
