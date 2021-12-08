package main

import (
	"log"
	"net/http"

	"github.com/Threx-code/go-blog/package/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.BlogRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9080", r))
}
