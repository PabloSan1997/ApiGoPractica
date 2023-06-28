package models

type Imagen struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}
type Imagenes []Imagen
