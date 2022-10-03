package app

import (
	"github.com/sharkx018/bookstore_items-api/controllers"
	"net/http"
)

func mapUrls() {
	router.HandleFunc("/item", controllers.ItemController.Create).Methods(http.MethodPost)
	router.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodGet)
}
