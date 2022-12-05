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
	pA3 := models.Preguntas{
		IdTest: "1",
		IdPregunta: "3",
		TypeTest: "autoestima",
		Pregunta: "Y cuando algo te sale muy bien!!?",
		Respuesta1:"chupame la pollaaa" ,
		Respuesta2: "kakakaka",
		Respuesta3: "lalalalala",
	}
	pA4 := models.Preguntas{
		IdTest: "1",
		IdPregunta: "4",
		TypeTest: "autoestima",
		Pregunta: "Cuando hablas con papá o mamá como te comportas?",
		Respuesta1:"ajjajaja" ,
		Respuesta2: "kakakaka",
		Respuesta3: "lalalalala", 
	}
	pA5 := models.Preguntas{
		IdTest: "1",
		IdPregunta: "5",
		TypeTest: "autoestima",
		Pregunta: "Cuando estas con un amigo que haceis?",
		Respuesta1:"ajjajaja" ,
		Respuesta2: "kakakaka",
		Respuesta3: "lalalalala",
	}


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
	config.DB.Create(&pA3)
	config.DB.Create(&pA4)
	config.DB.Create(&pA5)
	config.DB.Create(&pB1)
	config.DB.Create(&pB2)
	// config.DB.Create(&p5)

} 

func SetUsersExamples(w http.ResponseWriter, r *http.Request) {

	u1 := models.User{
		IdUser: "1",
		Nombre: "Eloy",
		Password: "1111",
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
		Password: "1234",
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


func SetDibujos(w http.ResponseWriter, r *http.Request){
	
	d1 := models.Dibujos{
		IdDibujo: 1,
		Tipo: "casa",
		Descripcion: "Tienes que dibujar una casa",
	}

	d2 := models.Dibujos{
		IdDibujo: 2,
		Tipo: "arbol",
		Descripcion: "Tienes que dibujar un arbol",
	}

	d3 := models.Dibujos{
		IdDibujo: 3,
		Tipo: "amigo",
		Descripcion: "Tienes que dibujar a un amigo",
	}

	d4 := models.Dibujos{
		IdDibujo: 4 ,
		Tipo: "familia",
		Descripcion: "Tienes que dibujar a tu familia",
	}

	d5 := models.Dibujos{
		IdDibujo: 5,
		Tipo:"perro",
		Descripcion: "Tienes que dibujar a un perro",
	}

	config.DB.Create(&d1)
	config.DB.Create(&d2)
	config.DB.Create(&d3)
	config.DB.Create(&d4)
	config.DB.Create(&d5)
}



func SetRewardsShop(w http.ResponseWriter, r *http.Request){
	
	d1 := models.RewardsShop{
		IdRewardShop:1,
		Name: "Cuadro",
		Price: 100,
		Url: "https://res.cloudinary.com/eloy411/image/upload/v1669653667/Duckdoc_xzayat.glb",
	}

	d2 := models.RewardsShop{
		IdRewardShop:2,
		Name: "Dinosaurio",
		Price: 200,
		Url: "https://res.cloudinary.com/eloy411/image/upload/v1669313425/yjusl94n4ufdrzeyd4kz.png",
	}

	d3 := models.RewardsShop{
		IdRewardShop:3,	
		Name: "Planta",
		Price: 100,
		Url: "https://res.cloudinary.com/eloy411/image/upload/v1669482452/usxqzllusv06jb29zoqp.png",
	}

	d4 := models.RewardsShop{
		IdRewardShop:4,
		Name: "Cohete",
		Price: 1000,
		Url: "https://res.cloudinary.com/eloy411/image/upload/v1669200885/hyy73nsz4cj82scqpo6h.jpg",
	}

	d5 := models.RewardsShop{
		IdRewardShop:5,
		Name: "Muñeco",
		Price: 300,
		Url: "https://res.cloudinary.com/eloy411/image/upload/v1669654515/rw3dvyozqlg9lwsi6zro.png",
	}

	d6 := models.RewardsShop{
		IdRewardShop:6,
		Name: "Moñigo",
		Price: 7000,
		Url: "https://res.cloudinary.com/eloy411/image/upload/v1669732993/fjwjnlnz2jxhtaegc36l.png",
	}

	config.DB.Create(&d1)
	config.DB.Create(&d2)
	config.DB.Create(&d3)
	config.DB.Create(&d4)
	config.DB.Create(&d5)
	config.DB.Create(&d6)

}


func SetRewardsUser(w http.ResponseWriter, r *http.Request){
	
	d1 := models.RewardsUsers{
		IdData : 1,
		Id_User: "1",
		Name: "Cuadro",
		IdRewardShop: "1",
		Url: "https://res.cloudinary.com/eloy411/image/upload/v1669653667/Duckdoc_xzayat.glb",
		
	}

	config.DB.Create(&d1)
}

func ExempleStateDucktor(w http.ResponseWriter, r *http.Request){
	
	d1 := models.StatesDucktorDaily{

		Id: 1,
		IdUser: "1",
		Date: time.Now(),
		Good: 0,
		Average: 0,
		Danger: 0,
	}

	config.DB.Create(&d1)
}