package elasticsearch

import (
	"context"
	"time"

	"github.com/olivere/elastic"
)

var (
	// Client the exported instance
	Client esClientInterface = &esClient{}
)

type esClientInterface interface {
	setClient(*elastic.Client)
	Index(string, interface{}) (*elastic.IndexResponse, error)
}

type esClient struct {
	client *elastic.Client
}

// Init initializes the client instance (should be invoked only one time)
func Init() {
	client, err := elastic.NewClient(
		elastic.SetURL("http://172.0.0.1:9200"),
		elastic.SetHealthcheckInterval(10*time.Second),
		// elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		// elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
	)
	if err != nil {
		panic(err)
	}

	Client.setClient(client)
}

func (c *esClient) setClient(client *elastic.Client) {
	c.client = client
}

func (c *esClient) Index(index string, doc interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	return c.client.Index().
		Index("items").
		BodyJson(doc).
		Do(ctx)
}