package models

type Preguntas struct {
	IdPregunta string `gorm:"primaryKey"`
	IdTest     string `gorm:"primaryKey"`
	TypeTest   string `json:"type_test"`	
	Pregunta   string `json:"pregunta"`
	Respuesta1 string `json:"respuesta1"`
	Respuesta2 string `json:"respuesta2"` 
	Respuesta3 string `json:"respuesta3"` 
	Respuesta4 string `json:"respuesta4,Default:null"` 
}


type GetPreguntas struct {
	IdTest int
}
  