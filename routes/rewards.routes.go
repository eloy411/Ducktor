package routes

import (
	"encoding/json"
	"log"
	"net/http"
)




func SaveCoins(w http.ResponseWriter, r *http.Request){

	/**CONFIG HEADERS*/
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	/**POST PARA GUARDAR LOS COINS */
	/**DEVUELVE FRASE*/

	/**CONFIG RESPONSE*/
	resp := make(map[string]string)
	resp["message"] = "coins almacenados"
	jsonResp, err := json.Marshal(resp)

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		return
	}


	w.Write(jsonResp)
}

func SaveRewards(w http.ResponseWriter, r *http.Request) {

	/**CONFIG HEADERS*/
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	/**POST PARA GUARDAR LA RECOMPENSA ADQUIRIDA */
	/**DEVUELVE FRASE*/

	resp := make(map[string]string)
	resp["message"] = "recompensas guardadas"
	jsonResp, err := json.Marshal(resp)

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		return
	}


	w.Write(jsonResp)
}