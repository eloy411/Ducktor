package routes
import (
	"net/http"

	"github.com/eloy411/project-M12-BACK/config"
	"github.com/eloy411/project-M12-BACK/models"
)

func SetConversationPreguntas(w http.ResponseWriter, r *http.Request) {

	pA1 := models.Preguntas{
		IdTest:     211,
		IdPregunta: "1",
		TypeTest:   "autoestima",
		Pregunta:   "Cuando te despiertas por la mañana...",
		Respuesta1: "Estoy triste de vez en cuando",
		Respuesta2: "Estoy triste muchas veces",
		Respuesta3: "Estoy triste siempre.",
	}
	pA2 := models.Preguntas{
		IdTest:     211,
		IdPregunta: "2",
		TypeTest:   "autoestima",
		Pregunta:   "Cuando quieres hacer algo...",
		Respuesta1: "Nunca me saldrá nada bien",
		Respuesta2: "No estoy seguro",
		Respuesta3: "Las cosas me saldrán bien",
	}
	pA3 := models.Preguntas{
		IdTest:     211,
		IdPregunta: "3",
		TypeTest:   "autoestima",
		Pregunta:   "Normalmente...",
		Respuesta1: "Hago bien la mayoría de las cosas",
		Respuesta2: "Hago mal muchas cosas!",
		Respuesta3: "Todo lo hago mal...",
	}
	pA4 := models.Preguntas{
		IdTest:     211,
		IdPregunta: "4",
		TypeTest:   "autoestima",
		Pregunta:   "De las cosas que me dices?",
		Respuesta1: "Me divierten muchas cosas",
		Respuesta2: "Me divierten algunas cosas",
		Respuesta3: "No me divierten nada",
	}
	pA5 := models.Preguntas{
		IdTest:     211,
		IdPregunta: "5",
		TypeTest:   "autoestima",
		Pregunta:   "Y como te sueles portar?",
		Respuesta1: "Soy malo siempre",
		Respuesta2: "Soy malo algunas veces",
		Respuesta3: "Soy malo muchas veces",
	}

	pB1 := models.Preguntas{
		IdTest:     212,
		IdPregunta: "1",
		TypeTest:   "autoestima",
		Pregunta:   "Te preocupa algo?",
		Respuesta1: "A veces pienso que me pueden pasar cosas malas",
		Respuesta2: "Me preocupa que me puedan pasar cosas malas",
		Respuesta3: "Estoy seguro de que me va a pasar cosas terribles",
	}

	pB2 := models.Preguntas{
		IdTest:     212,
		IdPregunta: "2",
		TypeTest:   "autoestima",
		Pregunta:   "Y de ti que me dices?",
		Respuesta1: "Me gusta como soy",
		Respuesta2: "No me gusta como soy",
		Respuesta3: "Me odio",
	}
	pB3 := models.Preguntas{
		IdTest:     212,
		IdPregunta: "3",
		TypeTest:   "autoestima",
		Pregunta:   "Cuando pasa algo malo...",
		Respuesta1: "Todas las cosas malas son culpa mía",
		Respuesta2: "Muchas cosas malas son culpa mía",
		Respuesta3: "Nada es culpa mía",
	}
	pB4 := models.Preguntas{
		IdTest:     212,
		IdPregunta: "4",
		TypeTest:   "autoestima",
		Pregunta:   "Piensas en morir?",
		Respuesta1: "No pienso morir",
		Respuesta2: "Me gustaría morir",
		Respuesta3: "Me gustaría morir pero no puedo",
	}
	pB5 := models.Preguntas{
		IdTest:     212,
		IdPregunta: "5",
		TypeTest:   "autoestima",
		Pregunta:   "A veces quieres llorar?",
		Respuesta1: "Todos los días quiero llorar",
		Respuesta2: "Tengo ganas muchos días",
		Respuesta3: "Muy pocas veces quiero llorar",
	}

	pC1 := models.Preguntas{
		IdTest:     213,
		IdPregunta: "1",
		TypeTest:   "autoestima",
		Pregunta:   "Te sueles preocupar por las cosas que tienes que hacer?",
		Respuesta1: "Las cosas me preocupan siempre",
		Respuesta2: "Las cosas me preocupan muchas veces",
		Respuesta3: "Las cosas me preocupan pocas veces",
	}

	pC2 := models.Preguntas{
		IdTest:     213,
		IdPregunta: "2",
		TypeTest:   "autoestima",
		Pregunta:   "Que piensas de la gente?",
		Respuesta1: "Me gusta estar con la gente",
		Respuesta2: "No me gusta estar con la gente muchas veces",
		Respuesta3: "No quiero estar con la gente nunca",
	}

	pC3 := models.Preguntas{
		IdTest:     213,
		IdPregunta: "3",
		TypeTest:   "autoestima",
		Pregunta:   "Cuando tienes que decidir algo...",
		Respuesta1: "No puedo decidirme",
		Respuesta2: "Me cuesta decidirme",
		Respuesta3: "Me decido fácilmente",
	}

	pC4 := models.Preguntas{
		IdTest:     213,
		IdPregunta: "4",
		TypeTest:   "autoestima",
		Pregunta:   "Como te ves a ti mismo?",
		Respuesta1: "Tengo buen aspecto",
		Respuesta2: "Hay algunas cosas que no me gustan de mi",
		Respuesta3: "Soy feo",
	}

	pC5 := models.Preguntas{
		IdTest:     213,
		IdPregunta: "5",
		TypeTest:   "autoestima",
		Pregunta:   "Te gusta hacer deberes?",
		Respuesta1: "Siempre me cuesta ponerme a hacer los deberes",
		Respuesta2: "Muchas veces me cuesta ponerme a hacer los deberes",
		Respuesta3: "No me cuesta nada ponerme a hacer los deberes",
	}

	pAAA1 := models.Preguntas{
		IdTest:     331,
		IdPregunta: "1",
		TypeTest:   "autoestima",
		Pregunta:   "Me siento tenso(a) o nervioso(a):",
		Respuesta1: "Todo el día",
		Respuesta2: "Gran parte del día",
		Respuesta3: "De vez en cuando",
		Respuesta4: "Nunca",
	}

	pAAA2 := models.Preguntas{
		IdTest:     331,
		IdPregunta: "2",
		TypeTest:   "autoestima",
		Pregunta:   "Sigo disfrutando de las cosas como siempre:",
		Respuesta1: "Igual que antes",
		Respuesta2: "No tanto como antes",
		Respuesta3: "Solamente un poco",
		Respuesta4: "Ya no disfruto como antes",
	}
	pAAA3 := models.Preguntas{
		IdTest:     331,
		IdPregunta: "3",
		TypeTest:   "autoestima",
		Pregunta:   "Siento una especie de temor como si algo malo fuera a suceder:",
		Respuesta1: "Sí y muy intenso",
		Respuesta2: "Sí, pero no muy intenso",
		Respuesta3: "Sí, pero no me preocupa",
		Respuesta4: "No siento nada de eso",
	}
	pAAA4 := models.Preguntas{
		IdTest:     331,
		IdPregunta: "4",
		TypeTest:   "autoestima",
		Pregunta:   "Soy capaz de reírme y ver el lado gracioso de las cosas:",
		Respuesta1: "Igual que siempre",
		Respuesta2: "Actualmente, algo menos",
		Respuesta3: "Actualmente, mucho menos",
		Respuesta4: "Actualmente, nada",
	}
	pAAA5 := models.Preguntas{
		IdTest:     331,
		IdPregunta: "5",
		TypeTest:   "autoestima",
		Pregunta:   "Tengo la cabeza llena de preocupaciones:",
		Respuesta1: "Todo el día",
		Respuesta2: "Gran parte del día",
		Respuesta3: "De vez en cuando",
		Respuesta4: "Nunca",
	}

	pAAB1 := models.Preguntas{
		IdTest:     332,
		IdPregunta: "1",
		TypeTest:   "autoestima",
		Pregunta:   "Me siento lento(a) y torpe:",
		Respuesta1: "Todo el día",
		Respuesta2: "Gran parte del día",
		Respuesta3: "De vez en cuando",
		Respuesta4: "Nunca",
	}
	pAAB2 := models.Preguntas{
		IdTest:     332,
		IdPregunta: "2",
		TypeTest:   "autoestima",
		Pregunta:   "Soy capaz de permanecer sentado(a) tranquilo(a) y relajado(a):",
		Respuesta1: "Siempre",
		Respuesta2: "A menudo",
		Respuesta3: "Raras veces",
		Respuesta4: "Nunca!",
	}
	pAAB3 := models.Preguntas{
		IdTest:     332,
		IdPregunta: "3",
		TypeTest:   "autoestima",
		Pregunta:   "He perdido el interés por mi aspecto personal:",
		Respuesta1: "Completamente",
		Respuesta2: "A menudo",
		Respuesta3: "Rara vez",
		Respuesta4: "Nada",
	}
	pAAB4 := models.Preguntas{
		IdTest:     332,
		IdPregunta: "4",
		TypeTest:   "autoestima",
		Pregunta:   "Experimento una desagradable sensación de “nervios y hormigueos” en el estómago:",
		Respuesta1: "Siempre",
		Respuesta2: "A menudo",
		Respuesta3: "Rara vez",
		Respuesta4: "Nunca",
	}
	pAAB5 := models.Preguntas{
		IdTest:     332,
		IdPregunta: "5",
		TypeTest:   "autoestima",
		Pregunta:   "Espero las cosas con ilusión:",
		Respuesta1: "Siempre",
		Respuesta2: "A menudo",
		Respuesta3: "Rara vez",
		Respuesta4: "Nunca",
	}

	pAAC1 := models.Preguntas{
		IdTest:     333,
		IdPregunta: "1",
		TypeTest:   "autoestima",
		Pregunta:   "Me siento inquieto(a) como si no pudiera parar de moverme:",
		Respuesta1: "Siempre",
		Respuesta2: "A menudo",
		Respuesta3: "Rara vez",
		Respuesta4: "Nunca",
	}
	pAAC2 := models.Preguntas{
		IdTest:     333,
		IdPregunta: "2",
		TypeTest:   "autoestima",
		Pregunta:   "Soy capaz de disfrutar con un buen libro o con un buen programa de radio o televisión:",
		Respuesta1: "Siempre",
		Respuesta2: "A menudo",
		Respuesta3: "Rara vez",
		Respuesta4: "Nunca",
	}
	pAAC3 := models.Preguntas{
		IdTest:     333,
		IdPregunta: "3",
		TypeTest:   "autoestima",
		Pregunta:   "Experimento de repente sensaciones de gran angustia o temor:",
		Respuesta1: "Siempre",
		Respuesta2: "A menudo",
		Respuesta3: "Rara vez",
		Respuesta4: "Nunca",
	}
	pAAC4 := models.Preguntas{
		IdTest:     333,
		IdPregunta: "4",
		TypeTest:   "autoestima",
		Pregunta:   "Cuentame, que pienas cuando alguien dice que eres una persona muy simpática, guapa o divertida??",
		Respuesta1: "Me gusta mucho!",
		Respuesta2: "Me da mucha verguenza!",
		Respuesta3: "Creo que me dicen mentiras!",
		Respuesta4: "Creo que me dicen mentiras!",
	}
	pAAC5 := models.Preguntas{
		IdTest:     333,
		IdPregunta: "5",
		TypeTest:   "autoestima",
		Pregunta:   "Cuentame, que pienas cuando alguien dice que eres una persona muy simpática, guapa o divertida??",
		Respuesta1: "Me gusta mucho!",
		Respuesta2: "Me da mucha verguenza!",
		Respuesta3: "Creo que me dicen mentiras!",
		Respuesta4: "Creo que me dicen mentiras!",
	}

	config.DB.Create(&pA1)
	config.DB.Create(&pA2)
	config.DB.Create(&pA3)
	config.DB.Create(&pA4)
	config.DB.Create(&pA5) 
	config.DB.Create(&pB1)
	config.DB.Create(&pB2)
	config.DB.Create(&pB3)
	config.DB.Create(&pB4)
	config.DB.Create(&pB5)
	config.DB.Create(&pC1)
	config.DB.Create(&pC2)
	config.DB.Create(&pC3)
	config.DB.Create(&pC4)
	config.DB.Create(&pC5)
	config.DB.Create(&pAAA1)
	config.DB.Create(&pAAA2)
	config.DB.Create(&pAAA3)
	config.DB.Create(&pAAA4)
	config.DB.Create(&pAAA5)
	config.DB.Create(&pAAB1)
	config.DB.Create(&pAAB2)
	config.DB.Create(&pAAB3)
	config.DB.Create(&pAAB4)
	config.DB.Create(&pAAB5)
	config.DB.Create(&pAAC1)
	config.DB.Create(&pAAC2)
	config.DB.Create(&pAAC3)
	config.DB.Create(&pAAC4)
	config.DB.Create(&pAAC5)
	// config.DB.Create(&p5)

}