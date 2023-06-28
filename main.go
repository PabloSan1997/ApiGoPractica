package main

import (
	"log"
	"net/http"

	"github.com/PabloSan1997/ApiGoPractica/handler"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", handler.IndexRoute).Methods("GET")
	router.HandleFunc("/api/v1/imagenes", handler.LeerTareas).Methods("GET")

	log.Fatal(http.ListenAndServe(":3001", router))
}
