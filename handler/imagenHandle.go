package handler

import (
	"encoding/json"
	"net/http"

	"github.com/PabloSan1997/ApiGoPractica/models"
)

var imagenes = models.Imagenes{
	{
		ID:   1,
		Name: "Universo",
		Url:  "https://concepto.de/wp-content/uploads/2014/08/universo-e1551279319307-800x400.jpg",
	},
}

func LeerTareas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(imagenes)
}
