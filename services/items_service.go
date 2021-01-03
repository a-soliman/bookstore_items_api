package services

import (
	"github.com/a-soliman/bookstore_items_api/domain/items"
	"github.com/a-soliman/bookstore_utils-go/rest_errors"
)

var (
	// ItemsService the exported instance
	ItemsService ItemsServiceInterface = &itemsService{}
)

// ItemsServiceInterface the items service interface
type ItemsServiceInterface interface {
	Get(string) (*items.Item, *rest_errors.RestErr)
	Create(items.Item) (*items.Item, *rest_errors.RestErr)
}

type itemsService struct{}

func (s *itemsService) Get(id string) (*items.Item, *rest_errors.RestErr) {
	return nil, nil
}

func (s *itemsService) Create(item items.Item) (*items.Item, *rest_errors.RestErr) {
	return nil, nil
}
