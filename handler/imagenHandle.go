package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/PabloSan1997/ApiGoPractica/models"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var imagenes = models.Imagenes{
	{
		ID:   uuid.New().String(),
		Name: "Universo",
		Url:  "https://concepto.de/wp-content/uploads/2014/08/universo-e1551279319307-800x400.jpg",
	},
	{
		ID:   uuid.New().String(),
		Name: "Universo",
		Url:  "https://concepto.de/wp-content/uploads/2014/08/universo-e1551279319307-800x400.jpg",
	},
	{
		ID:   uuid.New().String(),
		Name: "Universo",
		Url:  "https://concepto.de/wp-content/uploads/2014/08/universo-e1551279319307-800x400.jpg",
	},
}

func AgregarImagen(w http.ResponseWriter, r *http.Request) {
	var nuevaImagen models.Imagen
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Task Data")
	}
	json.Unmarshal(reqBody, &nuevaImagen)
	nuevaImagen.ID = uuid.New().String()
	imagenes = append(imagenes, nuevaImagen)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nuevaImagen)
}

func LeerImagenes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(imagenes)
}
func LeerUnaImagen(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	imagenesId := vars["id"]

	for _, ima := range imagenes {
		if ima.ID == imagenesId {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(ima)
		}
	}
}
func EliminarImagen(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	imagenId := vars["id"]
	var datosNuevos models.Imagenes
	for i, ima := range imagenes {
		if imagenId == ima.ID {
			datosNuevos = append(imagenes[:i], imagenes[i+1:]...)
			imagenes = datosNuevos
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusAccepted)
			json.NewEncoder(w).Encode(ima)
		}
	}

}

func EditarImagen(w http.ResponseWriter, r *http.Request) {

	var datosNuevos models.Imagen
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	json.Unmarshal(reqBody, &datosNuevos)

	for i, ima := range imagenes {
		if ima.ID == datosNuevos.ID {
			var nuevaData models.Imagenes = imagenes[:i]
			nuevaData = append(nuevaData, datosNuevos)
			nuevaData = append(nuevaData, imagenes[i+1:]...)
			imagenes = nuevaData
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusAccepted)
			json.NewEncoder(w).Encode(datosNuevos)
		}
	}
}
