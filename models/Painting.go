package models

import (

	"time"
)

type Painting struct {
	IdDibujo     uint64 `gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null"`
	NombreDibujo string
	Dibujo       string 
	IdUser       string `gorm:"FOREIGNKEY:IdUser,constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Fecha        time.Time `gorm:"autoCreateTime"`	
} 
