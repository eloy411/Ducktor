package models

import "time"

type User struct {
	IdUser       string `gorm:"primaryKey"`	
	Nombre       string
	Edad         int
	Ciudad       string
	Hospital     string
	Enfermedad   string
	Confianza    int
	Gravedad     int
	Numtest      int
	Numdibujos   int
	Fechaentrada time.Time `gorm:"autoCreateTime"`
}