package items

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sharkx018/bookstore_items-api/client/elasticsearch"
	"github.com/sharkx018/bookstore_items-api/domain/queries"
	"github.com/sharkx018/bookstore_utils-go/rest_errors"
	"log"
	"strings"
)

// DAO refers:  where we are going to store the item

const (
	indexItem = "items"
	indexType = "_doc"
)

func (i *Item) Save() rest_errors.RestErr {

	res, err := elasticsearch.Client.Index(indexItem, indexType, i)
	if err != nil {
		log.Println("==>>>>>> ", err)
		return rest_errors.NewInternalServerError("error when trying to save item", errors.New("database error"))
	}

	i.Id = res.Id
	return nil
}

func (i *Item) Get() rest_errors.RestErr {
	itemID := i.Id
	result, err := elasticsearch.Client.Get(indexItem, indexType, i.Id)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			return rest_errors.NewNotFoundError(fmt.Sprintf("no item found with id: %s", i.Id))
		}
		return rest_errors.NewInternalServerError(fmt.Sprintf("error when trying to get id %s", i.Id), errors.New("database error"))
	}

	bytess, _ := result.Source.MarshalJSON()

	json.Unmarshal(bytess, i)

	fmt.Println(result.Source)
	i.Id = itemID
	return nil
}

func (i *Item) Search(query queries.EsQuery) ([]Item, rest_errors.RestErr) {

	searchedResults, err := elasticsearch.Client.Search(indexItem, query.Build())
	if err != nil {
		return nil, rest_errors.NewInternalServerError("error when trying to search documents", errors.New("database error"))
	}
	fmt.Println("===>>>>>searchedResults ", searchedResults.TotalHits())

	items := make([]Item, searchedResults.TotalHits())
	for _, hits := range searchedResults.Hits.Hits {

		var item Item
		itemBytes, err := hits.Source.MarshalJSON()

		if err != nil {
			return nil, rest_errors.NewInternalServerError("error when trying to parse searched documents", errors.New("database error"))
		}

		err = json.Unmarshal(itemBytes, &item)
		if err != nil {
			return nil, rest_errors.NewInternalServerError("error when trying to parse searched documents", errors.New("database error"))
		}

		item.Id = hits.Id

		items = append(items, item)
	}

	return items, nil

}
