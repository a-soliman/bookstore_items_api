package items

import (
	"errors"

	"github.com/a-soliman/bookstore_items_api/clients/elasticsearch"
	"github.com/a-soliman/bookstore_utils-go/rest_errors"
)

const (
	indexItems = "items"
)

func (i *Item) Save() rest_errors.RestErr {
	result, err := elasticsearch.Client.Index(indexItems, i)
	if err != nil {
		return rest_errors.NewInternalServerError("error while trying to save item", errors.New("database error"))
	}
	i.ID = result.Id
	return nil
}
