package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PabloSan1997/ApiGoPractica/handler"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error con la variables de entorno")
	}
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = ":3001"
	}
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handler.IndexRoute).Methods("GET")
	router.HandleFunc("/api/v1/imagenes", handler.LeerImagenes).Methods("GET")
	router.HandleFunc("/api/v1/imagenes/{id}", handler.LeerUnaImagen).Methods("GET")
	router.HandleFunc("/api/v1/imagenes", handler.AgregarImagen).Methods("POST")
	router.HandleFunc("/api/v1/imagenes/{id}", handler.EliminarImagen).Methods("DELETE")
	router.HandleFunc("/api/v1/imagenes", handler.EditarImagen).Methods("PUT")
	log.Fatal(http.ListenAndServe(PORT, router))
}
