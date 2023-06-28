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
	router.HandleFunc("/api/v1/imagenes", handler.LeerImagenes).Methods("GET")
	router.HandleFunc("/api/v1/imagenes/{id}", handler.LeerUnaImagen).Methods("GET")
	router.HandleFunc("/api/v1/imagenes", handler.AgregarImagen).Methods("POST")
	router.HandleFunc("/api/v1/imagenes/{id}", handler.EliminarImagen).Methods("DELETE")
	router.HandleFunc("/api/v1/imagenes", handler.EditarImagen).Methods("PUT")
	log.Fatal(http.ListenAndServe(":3000", router))
}
