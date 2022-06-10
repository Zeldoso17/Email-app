package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"
	"github.com/gorilla/mux"
	"github.com/Zeldoso17/Email-App/routers"
)

func Managers() {
	router := mux.NewRouter()

	router.HandleFunc("/sendEmail", routers.SendEmail).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}