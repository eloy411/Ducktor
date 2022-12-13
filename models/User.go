package models

import "time"

type User struct {
	IdUser       string `gorm:"primaryKey"`	
	Nombre       string
	Password     string
	Edad         int
	Ciudad       string
	Hospital     string
	Enfermedad   string
	Confianza    int
	Gravedad     int
	Numtest      int
	Numdibujos   int
	Coins        int
	Fechaentrada time.Time `gorm:"autoCreateTime"`
}

type InfoUser struct {
	IdUser       string 
	Nombre       string 
	Edad         int
	Enfermedad   string
	Confianza    int
	Gravedad     int
	Numtest      int
	Numdibujos   int
	Coins        int
	Fechaentrada time.Time `gorm:"autoCreateTime"`
}

type Login struct {
	Nombre   string
	Password string
}