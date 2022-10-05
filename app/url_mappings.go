package app

import (
	"github.com/sharkx018/bookstore_items-api/controllers"
	"net/http"
)

func mapUrls() {

	router.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodGet)

	router.HandleFunc("/item", controllers.ItemController.Create).Methods(http.MethodPost)
	router.HandleFunc("/items/{id}", controllers.ItemController.Get).Methods(http.MethodGet)
	router.HandleFunc("/items/search", controllers.ItemController.Search).Methods(http.MethodPost)
}
