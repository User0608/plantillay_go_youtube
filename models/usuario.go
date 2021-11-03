package models

type RequestSession struct {
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
}

type Usuario struct {
	Nombre   string `json:"nombre"`
	Edad     int    `json:"edad"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
}
