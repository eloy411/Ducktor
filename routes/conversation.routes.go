package routes

import (
	"encoding/json"
	
	"log"
	"net/http"

	"github.com/eloy411/project-M12-BACK/models"
)

type Preguntas struct {
	Pregunta1 string
	Pregunta2 string
	Pregunta3 string
	Pregunta4 string
	Pregunta5 string
	Pregunta6 string
}








func GetConversation(w http.ResponseWriter, r *http.Request) {
	/**CONFIG HEADERS*/
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")



	p := Preguntas{
		Pregunta1: "soy la pregunta 1",
		Pregunta2: "soy la pregunta 2",
		Pregunta3: "soy la pregunta 3",
		Pregunta4: "soy la pregunta 4",
		Pregunta5: "soy la pregunta 5",
		Pregunta6: "soy la pregunta 6",
	}
	/** REQUEST PARA PEDIR LAS CONVERSACIONES SEGUN DIA, CONFIANZA, NIVEL DE ENFERMEDAD, NUM CONVERSACIONES*/
	resp := make(map[string]string)
	resp["message"] = "iniciando conversacion"
	jsonResp, err := json.Marshal(p)




	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		return
	}
	w.Write(jsonResp)

	
}


func RegisterResponses(w http.ResponseWriter, r *http.Request) {

	/**CONFIG HEADERS*/
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	/** REGISTRAR LAS RESPUESTAS */

	var respuestas models.Responses

	err := json.NewDecoder(r.Body).Decode(&respuestas)

 	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	w.Write([]byte(respuestas.Respuesta1))


	/**RESPONDER AL CLIENTE*/

	
	resp := make(map[string]string)
	resp["message"] = "almacenando respuestas"
	jsonResp, err := json.Marshal(resp)

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		return
	}


	w.Write(jsonResp)
}