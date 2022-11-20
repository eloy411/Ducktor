package models

import "time"

type User struct {
	IdUser       string
	Nombre       string
	Edad         int
	Ciudad       string
	Hospital     string
	Enfermedad   string
	Confianza    string
	Gravedad     string
	Numtest      int
	Numdibujos   string
	Fechaentrada time.Time
}