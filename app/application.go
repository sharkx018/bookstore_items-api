package app

import (
	"github.com/gorilla/mux"
	"github.com/sharkx018/bookstore_items-api/client/elasticsearch"
	"log"
	"net/http"
	"time"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {

	elasticsearch.Init()
	mapUrls()

	srv := &http.Server{
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 500 * time.Millisecond,
		ReadTimeout:  2 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      router,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start the server %s", err)
	}
}
