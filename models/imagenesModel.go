package models

type Imagen struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}
type Imagenes []Imagen
