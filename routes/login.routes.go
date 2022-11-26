package routes

import (
	"encoding/json"
	"fmt"

	"log"
	"net/http"

	"github.com/eloy411/project-M12-BACK/config"
	"github.com/eloy411/project-M12-BACK/models"
)

func Login(w http.ResponseWriter, r *http.Request) {

	/**CONFIG HEADERS*/
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	/**COMPROBAMOS SI EL USUARIO EXISTE EN LA BASE DE DATOS*/
	
	var login models.Login

	err := json.NewDecoder(r.Body).Decode(&login)
	fmt.Println(login)
	user := models.InfoUser{}
	
	config.DB.Table("users").Select("*").Where("Nombre = ? AND password = ?", &login.Nombre, &login.Password).Scan(&user)
	
	fmt.Println(user)
	if user.Nombre == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "Usuario no encontrado"}`))
		return
	}

	
	jsonResp, err := json.Marshal(&user)

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		return
	}

	w.Write(jsonResp)
	
}