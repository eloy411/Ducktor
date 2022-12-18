package models

import (
	"time"
)


type Dibujos struct {
	IdDibujo    int
	Tipo        string
	Descripcion string
}

type FileData struct { 
	MyFile string
}

type GetDibujo struct { 
	IdDibujo int
}

type PaintingDaily struct { 
	IdPaintingDaily int `gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null"`
	IdDibujo     string
	NombreDibujo string
	URL_Dibujo   string
	IdUser       string    `gorm:"PRIMARY_KEY"`
	Fecha        time.Time `gorm:"autoCreateTime;PRIMARY_KEY"`
} 

 
