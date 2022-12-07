package models

import "time"

type ResponsesTestsDaily struct {
	Id_response int    `gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null"`
	IdUser      string `gorm:"FOREINGKEY:IdUser,constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	IdTest      string `gorm:"FOREINGKEY:IdTest,constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Respuesta1  int
	Respuesta2  int
	Respuesta3  int
	Respuesta4  int
	Respuesta5  int
	Fecha       time.Time `gorm:"autoCreateTime"`  
}  
