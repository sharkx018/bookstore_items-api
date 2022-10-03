package items

import (
	"errors"
	"github.com/sharkx018/bookstore_items-api/client/elasticsearch"
	"github.com/sharkx018/bookstore_utils-go/rest_errors"
)

// DAO refers:  where we are going to store the item

const (
	indexItem = "items"
)

func (i *Item) Save() rest_errors.RestErr {
	res, err := elasticsearch.Client.Index(indexItem, i)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to save item", errors.New("database error"))
	}

	i.Id = res.Id
	return nil
}
