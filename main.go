package main

import (
	"net/http"

	"github.com/eloy411/project-M12-BACK/config"
	"github.com/eloy411/project-M12-BACK/models"
	"github.com/eloy411/project-M12-BACK/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	
	config.ConnectDB()
	config.DB.AutoMigrate(&models.Preguntas{})
	config.DB.AutoMigrate(&models.User{})
	config.DB.AutoMigrate(&models.ResponsesTestsDaily{})
	config.DB.AutoMigrate(&models.Dibujos{})
	config.DB.AutoMigrate(&models.PaintingDaily{})
	config.DB.AutoMigrate(&models.RewardsShop{})
	config.DB.AutoMigrate(&models.RewardsUsers{})
	config.DB.AutoMigrate(&models.StatesDucktorDaily{})
	
	config.ConncetCloudinary()
	r := mux.NewRouter()

	/**ROUTES*/

	r.HandleFunc("/login", routes.Login).Methods("POST")

	r.HandleFunc("/ini",routes.Initialization).Methods("POST")
	r.HandleFunc("/finish",routes.Finish).Methods("POST")
 
	r.HandleFunc("/game1",routes.IniGame).Methods("POST")
	r.HandleFunc("/game1",routes.DataGame).Methods("POST")

	r.HandleFunc("/conversation-init",routes.GetConversation).Methods("POST")
	r.HandleFunc("/conversation-register-data",routes.RegisterResponses).Methods("POST")

	r.HandleFunc("/painting-init",routes.IniPainting).Methods("POST")
	r.HandleFunc("/save-painting",routes.SavePaint).Methods("POST")

	r.HandleFunc("/rewards-shop",routes.SendRewardsShop).Methods("GET")
	r.HandleFunc("/rewards-coins",routes.SaveCoins).Methods("PUT")
	r.HandleFunc("/rewards-user",routes.SaveRewardsUser).Methods("POST")
	r.HandleFunc("/send-rewards-user",routes.GetRewardsUser).Methods("POST")


	r.HandleFunc("/states-ducktor",routes.SetStatesDucktor).Methods("POST")
	r.HandleFunc("/states-ducktor",routes.UpdateStatesDucktor).Methods("PUT")
 
	r.HandleFunc("/admin-preguntas",routes.SetConversationPreguntas).Methods("GET")
	r.HandleFunc("/admin-users",routes.SetUsersExamples).Methods("GET")
	r.HandleFunc("/admin-dibujos",routes.SetDibujos)
	r.HandleFunc("/admin-rewards-shop",routes.SetRewardsShop).Methods("GET")
	r.HandleFunc("/admin-rewards-user",routes.SetRewardsUser).Methods("GET")

	handler := cors.Default().Handler(r)

	http.ListenAndServe(":3000",handler)
}