package main

import (
	"net/http"

	"github.com/eloy411/project-M12-BACK/routes"
	"github.com/rs/cors"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	/**ROUTES*/
	r.HandleFunc("/ini",routes.Initialization).Methods("GET")
	r.HandleFunc("/finish",routes.Finish).Methods("GET")

	r.HandleFunc("/game1",routes.IniGame).Methods("POST")
	r.HandleFunc("/game1",routes.DataGame).Methods("POST")

	r.HandleFunc("/conversation-init",routes.GetConversation).Methods("POST")
	r.HandleFunc("/conversation-register-data",routes.RegisterResponses).Methods("POST")

	r.HandleFunc("/painting-init",routes.IniPainting).Methods("POST")
	r.HandleFunc("/painting",routes.SavePaint).Methods("POST")

	r.HandleFunc("/rewards",routes.SaveCoins).Methods("PUT")
	r.HandleFunc("/rewards",routes.SaveRewards).Methods("POST")
 

	handler := cors.Default().Handler(r)

	http.ListenAndServe(":3000",handler)
}