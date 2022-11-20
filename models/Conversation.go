package models

type Preguntas struct {
	IdPregunta string `gorm:"primaryKey"`
	IdTest     string `gorm:"primaryKey"`
	TypeTest   string
	Pregunta   string
	Respuesta1 string
	Respuesta2 string
	Respuesta3 string
	// Response  []Responses `gorm:"ForeignKey:IdPregunta"`
}
// type Responses struct {
// 	IdRespuesta string `gorm:"primaryKey"`
// 	Respuesta   string 
// 	IdPregunta  string 
// 	Preguntas   Preguntas `gorm:"foreignKey:IdPregunta"`
// }
 