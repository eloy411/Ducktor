package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// "strconv"

	"github.com/eloy411/project-M12-BACK/config"
	"github.com/eloy411/project-M12-BACK/models"
)




func GetConversation(w http.ResponseWriter, r *http.Request) {
	/**CONFIG HEADERS*/
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	/** RECIBIMOS LOS DATOS DEL USUARIO*/

	var user models.GetPreguntas
	
	
	err := json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(user.IdTest)
	/** REQUEST PARA PEDIR LAS CONVERSACIONES SEGUN DIA, CONFIANZA, NIVEL DE ENFERMEDAD, NUM CONVERSACIONES*/
	var result []models.Preguntas
	config.DB.Table("preguntas").Select("*").Where("Id_Test = ?",user.IdTest+1).Scan(&result)

	jsonResp, err := json.Marshal(&result)

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		return
	}

	/**RESPONSE*/
	w.Write(jsonResp)

	
}


func RegisterResponses(w http.ResponseWriter, r *http.Request) {

	/**CONFIG HEADERS*/
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	/** REGISTRAR LAS RESPUESTAS */

	var respuestas models.ResponsesTestsDaily
	var user models.User


	err := json.NewDecoder(r.Body).Decode(&respuestas)
	fmt.Println(respuestas)
 	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	config.DB.Create(&respuestas)
	config.DB.Model(&user).Where("Id_User = ?",respuestas.IdUser).Update("NumTest",respuestas.IdTest)


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