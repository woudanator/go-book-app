package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/woudanator/go-book-app/pkg/routes"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8000", router))
}
