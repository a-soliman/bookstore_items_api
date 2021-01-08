package items

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/a-soliman/bookstore_items_api/clients/elasticsearch"
	"github.com/a-soliman/bookstore_utils-go/rest_errors"
)

const (
	indexItems  = "items"
	docTypeItem = "item"
)

func (i *Item) Save() rest_errors.RestErr {
	result, err := elasticsearch.Client.Index(indexItems, docTypeItem, i)
	if err != nil {
		return rest_errors.NewInternalServerError("error while trying to save item", errors.New("database error"))
	}
	i.ID = result.Id
	return nil
}

func (i *Item) Get() rest_errors.RestErr {
	itemId := i.ID
	result, err := elasticsearch.Client.Get(indexItems, docTypeItem, i.ID)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			return rest_errors.NewNotFoundError(fmt.Sprintf("no item found with id %s", i.ID))
		}
		return rest_errors.NewInternalServerError("error while trying to get item", errors.New("database error"))
	}
	if result == nil {
		return rest_errors.NewNotFoundError(fmt.Sprintf("no item found with id %s", i.ID))
	}

	bytes, err := result.Source.MarshalJSON()
	if err != nil {
		return rest_errors.NewInternalServerError("error while trying to parse database response", errors.New("database error"))
	}

	if err := json.Unmarshal(bytes, &i); err != nil {
		return rest_errors.NewInternalServerError("error while trying to parse database response", errors.New("database error"))
	}
	i.ID = itemId
	return nil
}
