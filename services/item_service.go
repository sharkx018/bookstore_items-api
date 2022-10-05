package services

import (
	"github.com/sharkx018/bookstore_items-api/domain/items"
	"github.com/sharkx018/bookstore_items-api/domain/queries"
	"github.com/sharkx018/bookstore_utils-go/rest_errors"
)

var (
	ItemService itemServiceInterface = &itemService{}
)

type itemServiceInterface interface {
	Create(item items.Item) (*items.Item, rest_errors.RestErr)
	Get(string) (*items.Item, rest_errors.RestErr)
	Search(query queries.EsQuery) ([]items.Item, rest_errors.RestErr)
}

type itemService struct{}

func (s *itemService) Create(item items.Item) (*items.Item, rest_errors.RestErr) {

	if err := item.Save(); err != nil {
		return nil, err
	}

	return &item, nil

}

func (s *itemService) Get(id string) (*items.Item, rest_errors.RestErr) {
	item := items.Item{
		Id: id,
	}

	if err := item.Get(); err != nil {
		return nil, err
	}

	return &item, nil

	//return nil, rest_errors.NewRestError("Implement me!", http.StatusNotImplemented, "not_implemented", nil)
}

func (s *itemService) Search(query queries.EsQuery) ([]items.Item, rest_errors.RestErr) {
	dao := items.Item{}
	items, err := dao.Search(query)
	if err != nil {
		return nil, err
	}
	return items, nil
}
