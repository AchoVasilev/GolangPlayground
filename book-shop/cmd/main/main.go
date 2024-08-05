package main

import (
	"github.com/achovasilev/book-shop/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)

	http.Handle("/", router)
	log.Println("Starting http server on port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
