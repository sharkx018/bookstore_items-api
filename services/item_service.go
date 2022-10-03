package services

import (
	"github.com/sharkx018/bookstore_items-api/domain/items"
	"github.com/sharkx018/bookstore_utils-go/rest_errors"
	"net/http"
)

var (
	ItemService itemServiceInterface = &itemService{}
)

type itemServiceInterface interface {
	Create(item items.Item) (*items.Item, rest_errors.RestErr)
	Get(string) (*items.Item, rest_errors.RestErr)
}

type itemService struct{}

func (s *itemService) Create(item items.Item) (*items.Item, rest_errors.RestErr) {

	if err := item.Save(); err != nil {
		return nil, err
	}

	return &item, nil

}

func (s *itemService) Get(string) (*items.Item, rest_errors.RestErr) {
	return nil, rest_errors.NewRestError("Implement me!", http.StatusNotImplemented, "not_implemented", nil)
}
