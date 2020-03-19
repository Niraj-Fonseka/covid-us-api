package main

import (
	"covid-us-api/handlers"
	"covid-us-api/services"
	"fmt"
	"log"
	"net/http"
)

func main() {

	log.Println("Starting server ...")
	services := services.RegisterServices()
	handlers := handlers.RegisterHandlers(services)

	http.HandleFunc("/daily", handlers.SlackHandler)
	http.HandleFunc("/draw", handlers.DrawGraph)
	http.HandleFunc("/drawstate", handlers.DrawGraphState)

	http.HandleFunc("/test", handlers.DrawGraphUSMAP)
	fmt.Println("Listening..")
	http.ListenAndServe(":8080", nil)

}
