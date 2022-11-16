package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/eloy411/project-M12-BACK/models"
)




func GetConversation(w http.ResponseWriter, r *http.Request) {
	/**CONFIG HEADERS*/
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	/** RECIBIMOS LOS DATOS DEL USUARIO*/

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	/**ESTA SIMULANDO UNA PETICION A LA BASE DE DATOS hacer una query que sacara preguntas según la que no haya hecho, gravedad, tiempo de estancia */
	p := models.Preguntas{
		Pregunta1: "soy la pregunta 1",
		Pregunta2: "soy la pregunta 2",
		Pregunta3: "soy la pregunta 3",
		Pregunta4: "soy la pregunta 4",
		Pregunta5: "soy la pregunta 5",
	}

	/** REQUEST PARA PEDIR LAS CONVERSACIONES SEGUN DIA, CONFIANZA, NIVEL DE ENFERMEDAD, NUM CONVERSACIONES*/
	
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