package models

import "time"

type User struct {
	IdUser       string
	Nombre       string
	Edad         int
	Ciudad       string
	Hospital     string
	Enfermedad   string
	Confianza    int
	Gravedad     int
	FechaEntrada time.Time
	NumTest      int
	NumDibujos   int
}