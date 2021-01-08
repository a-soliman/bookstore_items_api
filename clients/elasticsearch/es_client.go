package elasticsearch

import (
	"context"
	"fmt"
	"time"

	"github.com/a-soliman/bookstore_utils-go/logger"
	"github.com/olivere/elastic"
)

var (
	// Client the exported instance
	Client esClientInterface = &esClient{}
)

type esClientInterface interface {
	setClient(*elastic.Client)
	Index(string, string, interface{}) (*elastic.IndexResponse, error)
	Get(string, string, string) (*elastic.GetResult, error)
	Search(string, string, elastic.Query) (*elastic.SearchResult, error)
}

type esClient struct {
	client *elastic.Client
}

// Init initializes the client instance (should be invoked only one time)
func Init() {
	log := logger.GetLogger()
	client, err := elastic.NewClient(
		elastic.SetURL("http://localhost:9200"),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetErrorLog(log),
		elastic.SetInfoLog(log),
	)
	if err != nil {
		panic(err)
	}

	Client.setClient(client)
}

func (c *esClient) setClient(client *elastic.Client) {
	c.client = client
}

func (c *esClient) Index(index string, docType string, doc interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	result, err := c.client.Index().
		Index(index).
		Type(docType).
		BodyJson(doc).
		Do(ctx)

	if err != nil {
		// log the error
		return nil, err
	}
	return result, nil
}

func (c *esClient) Get(index string, docType string, documentID string) (*elastic.GetResult, error) {
	ctx := context.Background()
	result, err := c.client.Get().
		Index(index).
		Type(docType).
		Id(documentID).
		Do(ctx)

	if err != nil {
		logger.Error(fmt.Sprintf("error while trying to get id %s", documentID), err)
		return nil, err
	}
	if !result.Found {
		return nil, nil
	}
	return result, nil
}

func (c *esClient) Search(index string, docType string, query elastic.Query) (*elastic.SearchResult, error) {
	ctx := context.Background()
	result, err := c.client.Search(index).
		Query(query).
		RestTotalHitsAsInt(true).
		Do(ctx)

	if err != nil {
		logger.Error(fmt.Sprintf("error while trying to search documents in index %s", index), err)
		return nil, err
	}
	return result, nil
}
