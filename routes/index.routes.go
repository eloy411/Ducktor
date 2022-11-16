package routes

import (
	// "encoding/json"
	// "fmt"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Person struct {
    Name string
    Age  string
}

func Initialization(w http.ResponseWriter, r *http.Request) {

    /**CONFIG HEADERS*/
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

    /**CARGA LA INFORMACIÓN PERSONAL DEL NIÑO*/
    /**CARGAR LA INFORMACION DEL NIÑO ENTORNO AL JUEGO (COINS PARA INTERCAMBIAR RECOMPENSAS,ETC)*/
	resp := make(map[string]string)
	resp["message"] = "Has iniciado el juego"
	jsonResp, err := json.Marshal(resp)

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		return
	}


	w.Write(jsonResp)
}

func Finish(w http.ResponseWriter, r *http.Request) {

    // /**CONFIG HEADERS*/
	// w.WriteHeader(http.StatusCreated)
	// w.Header().Set("Content-Type", "application/json")

    // /** GUARDAMOS LA INFORMACIÓN RECOPILADA DESDE EL CLIENTE (EN EL CACHÉ) Y LA ALMACENAMOS EN LA TABLA "DATOS DIARIOS" */

	// resp := make(map[string]string)
	// resp["message"] = "Se esta finalizando el juego"
	// jsonResp, err := json.Marshal(resp)

	// if err != nil {
	// 	log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	// 	return
	// }


	// w.Write(jsonResp)
// }

	var p Person

	err := json.NewDecoder(r.Body).Decode(&p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Do something with the Person struct...
    fmt.Fprintf(w, "Person: %+v", p)
}