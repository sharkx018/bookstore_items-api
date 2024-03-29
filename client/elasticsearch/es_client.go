package elasticsearch

import (
	"fmt"
	//"fmt"
	"github.com/olivere/elastic"
	"github.com/sharkx018/bookstore_utils-go/logger"
	"golang.org/x/net/context"
	_ "log"
	_ "os"
	"time"
)

var (
	Client esClientInterface = &esClient{}
)

type esClient struct {
	client *elastic.Client
}

type esClientInterface interface {
	setClient(*elastic.Client)
	Index(string, string, interface{}) (*elastic.IndexResponse, error)
	Get(index string, docType string, id string) (*elastic.GetResult, error)
	Search(index string, query elastic.Query) (*elastic.SearchResult, error)
}

func Init() {

	log := logger.GetLogger()
	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetErrorLog(log),
		elastic.SetInfoLog(log),
		//elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		//elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
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
	result, err := c.client.Index().Index(index).Type(docType).BodyJson(doc).Do(ctx)

	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to index document in index, %s", index), err)
		return nil, err
	}

	return result, nil
}

func (c *esClient) Get(index string, docType string, id string) (*elastic.GetResult, error) {
	ctx := context.Background()
	result, err := c.client.Get().
		Index(index).
		Type(docType).
		Id(id).
		Do(ctx)

	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to get id %s", id), err)
		return nil, err
	}

	return result, nil

}

func (c *esClient) Search(index string, query elastic.Query) (*elastic.SearchResult, error) {
	ctx := context.Background()
	result, err := c.client.Search(index).Query(query).RestTotalHitsAsInt(true).Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to search documents in index %s", query), err)
		return nil, err
	}
	return result, nil
}
