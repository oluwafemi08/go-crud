package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_"github.com/lib/pq"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/oluwafemi08/go-app/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterAppRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost: 9010", r)) //localhost: 9010

}
