package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sharkx018/bookstore_items-api/domain/items"
	"github.com/sharkx018/bookstore_items-api/domain/queries"
	"github.com/sharkx018/bookstore_items-api/services"
	"github.com/sharkx018/bookstore_items-api/utils/http_utils"
	"github.com/sharkx018/bookstore_oauth-go/oauth"
	"github.com/sharkx018/bookstore_utils-go/rest_errors"
	"io"
	"io/ioutil"
	"net/http"
)

var (
	ItemController itemControllerInterface = &itemController{}
)

type itemControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Search(w http.ResponseWriter, r *http.Request)
}

type itemController struct{}

func (c *itemController) Create(w http.ResponseWriter, r *http.Request) {

	errR := oauth.AuthenticationRequest(r)
	if errR != nil {
		//http_utils.RespondError(w, err)
		// TODO it is placed just to return some error
		restErr := rest_errors.NewUnauthorizedError("invalid request body")
		http_utils.RespondError(w, restErr)
		return
	}

	sellerId := oauth.GetCallerId(r)
	if sellerId == 0 {
		restErr := rest_errors.NewUnauthorizedError("invalid request body")
		http_utils.RespondError(w, restErr)
		return
	}

	//read the body into bytes
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		restErr := rest_errors.NewBadRequestError("invalid request body")
		http_utils.RespondError(w, restErr)
		return
	}
	defer r.Body.Close()

	var item items.Item
	// take the bytes and cast to struct
	err = json.Unmarshal(requestBody, &item)
	if err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.RespondError(w, restErr)
		return
	}

	// set the oauth caller id
	item.Seller = sellerId

	result, createErr := services.ItemService.Create(item)
	if createErr != nil {
		http_utils.RespondError(w, createErr)
		return
	}

	fmt.Println(result)
	http_utils.RespondJson(w, http.StatusCreated, result)

}

func (c *itemController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemID := vars["id"]
	item, err := services.ItemService.Get(itemID)
	if err != nil {
		http_utils.RespondError(w, err)
		return
	}

	http_utils.RespondJson(w, http.StatusOK, item)

}

func (c *itemController) Search(w http.ResponseWriter, r *http.Request) {

	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.RespondError(w, restErr)
		return
	}
	defer r.Body.Close()

	var query queries.EsQuery
	err = json.Unmarshal(reqBytes, &query)
	if err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.RespondError(w, restErr)
		return
	}

	items, searchErr := services.ItemService.Search(query)
	if searchErr != nil {
		http_utils.RespondError(w, searchErr)
		return
	}

	http_utils.RespondJson(w, http.StatusOK, items)

}
