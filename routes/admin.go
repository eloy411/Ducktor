package routes

import (
	"net/http"
	"time"

	"github.com/eloy411/project-M12-BACK/config"
	"github.com/eloy411/project-M12-BACK/models"
)

func SetConversationPreguntas(w http.ResponseWriter, r *http.Request) {

	pA1 := models.Preguntas{
		IdTest: "1",
		IdPregunta: "1",
		TypeTest: "autoestima",
		Pregunta: "Cuando te levantas por la mañana que es lo que mas quieres hacer??",
		Respuesta1:"Volver a dormir!" ,
		Respuesta2:"Ver a mi mamá o a mi papá!" ,
		Respuesta3: "Jugar a algo!",
	}
	pA2 := models.Preguntas{
		IdTest: "1",
		IdPregunta: "2",
		TypeTest: "autoestima",
		Pregunta: "Cuando quieres hacer algo y te sale mal que sueles hacer?",
		Respuesta1:"Nada" ,
		Respuesta2:"Enfadarme mucho" ,
		Respuesta3: "Decir que no pasa nada y que volveré a intentarlo",
	}
	// pA3 := models.Preguntas{
	// 	IdTest: "1",
	// 	IdPregunta: "3",
	// 	TypeTest: "autoestima",
	// 	Pregunta: "Y cuando algo te sale muy bien!!?",
	// 	Respuesta1: ,
	// 	Respuesta2: ,
	// 	Respuesta3: 
	// }
	// pA4 := models.Preguntas{
	// 	IdTest: "1",
	// 	IdPregunta: "4",
	// 	TypeTest: "autoestima",
	// 	Pregunta: "Cuando hablas con papá o mamá como te comportas?",
	// 	Respuesta1: ,
	// 	Respuesta2: ,
	// 	Respuesta3: 
	// }
	// pA5 := models.Preguntas{
	// 	IdTest: "1",
	// 	IdPregunta: "5",
	// 	TypeTest: "autoestima",
	// 	Pregunta: "Cuando estas con un amigo que haceis?",
	// 	Respuesta1: ,
	// 	Respuesta2: ,
	// 	Respuesta3: 
	// }


	pB1 := models.Preguntas{
		IdTest: "2",
		IdPregunta: "1",
		TypeTest: "autoestima",
		Pregunta: "Si alguien que es tu amigo, de repente hace algo que no te gusta?",
		Respuesta1:"Me pone muy triste!" ,
		Respuesta2: "Me da igual!",
		Respuesta3: "Me enfado mucho",
	}

	pB2 := models.Preguntas{
		IdTest: "2",
		IdPregunta: "2",
		TypeTest: "autoestima",
		Pregunta: "Cuentame, que pienas cuando alguien dice que eres una persona muy simpática, guapa o divertida??",
		Respuesta1:"Me gusta mucho!" ,
		Respuesta2: "Me da mucha verguenza!",
		Respuesta3: "Creo que me dicen mentiras!",
	}
	config.DB.Create(&pA1)
	config.DB.Create(&pA2)
	config.DB.Create(&pB1)
	config.DB.Create(&pB2)
	// config.DB.Create(&p5)

} 

func SetUsersExamples(w http.ResponseWriter, r *http.Request) {

	u1 := models.User{
		IdUser: "1",
		Nombre: "Eloy",
		Edad: 12,
		Ciudad: "Barcelona",
		Hospital: "Hospital de Sant Pau",
		Enfermedad: "Cancer",
		Confianza: 0,
		Gravedad: 0,
		Numtest: 0,
		Numdibujos: 0,
		Fechaentrada: time.Now(),
	}
	u2 := models.User{
		IdUser: "2",
		Nombre: "Pablo",
		Edad: 8,
		Ciudad: "Barcelona",
		Hospital: "Hospital de Sant Pau",
		Enfermedad: "Rotura de tibia",
		Confianza: 0,
		Gravedad: 0,
		Numtest: 0,
		Numdibujos: 0,
		Fechaentrada: time.Now(),
	}

	config.DB.Create(&u1)
	config.DB.Create(&u2)
}