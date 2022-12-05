package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/eloy411/project-M12-BACK/config"
	"github.com/eloy411/project-M12-BACK/models"
)

func SetStatesDucktor(w http.ResponseWriter, r *http.Request) {

	var states models.StatesDucktorDaily
	
	err := json.NewDecoder(r.Body).Decode(&states)
	
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		return
	}
	config.DB.Create(&states)

}


func UpdateStatesDucktor(w http.ResponseWriter, r *http.Request) {

	var states models.StatesDucktorDaily
	
	err := json.NewDecoder(r.Body).Decode(&states)
	
	fmt.Println(states)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		return
	}
	config.DB.Model(&states).Where("Id_User = ?", states.IdUser).Updates(models.StatesDucktorDaily{Good: states.Good, Average: states.Average, Danger: states.Danger})
}