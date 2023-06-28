package handler

import (
	"encoding/json"
	"net/http"
)

type message struct {
	Mes string `json:"mensaje"`
}

func IndexRoute(w http.ResponseWriter, r *http.Request) {

	mensaje := message{Mes: "Bienvenido a mi api"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mensaje)
}
