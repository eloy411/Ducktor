package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/eloy411/project-M12-BACK/config"
	"github.com/eloy411/project-M12-BACK/models"
)


func SendRewardsShop(w http.ResponseWriter, r *http.Request) {

	var rewards []models.RewardsShop;

	config.DB.Table("rewards_shop").Select("*").Scan(&rewards)

	jsonResp, err := json.Marshal(&rewards)

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		return
	}

	/**RESPONSE*/
	w.Write(jsonResp)

}


func SaveCoins(w http.ResponseWriter, r *http.Request){

	/**CONFIG HEADERS*/
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	/**POST PARA GUARDAR LOS COINS */
	var coins models.Coins
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&coins)

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		return
	}
	config.DB.Model(&user).Where("Id_User = ?", coins.IdUser).Update("Coins", coins.Coins)
	
}



func SaveRewardsUser(w http.ResponseWriter, r *http.Request) {

	var rewardsUser models.RewardsUsers

	err := json.NewDecoder(r.Body).Decode(&rewardsUser)

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		return
	}

	config.DB.Create(&rewardsUser)

}

func GetRewardsUser(w http.ResponseWriter, r *http.Request) {

	var user models.InfoUser

	
	var rewardsUser []models.RewardsUsers

	config.DB.Table("rewards_users").Select("*").Where("Id_User = ?",user.IdUser).Scan(&rewardsUser)

	jsonResp, err := json.Marshal(&rewardsUser)

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		return
	}

	/**RESPONSE*/
	w.Write(jsonResp)

}