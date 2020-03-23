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
	svcs := services.RegisterServices()
	pages := services.NewPages(svcs)
	handlers := handlers.RegisterHandlers(svcs, pages)

	http.HandleFunc("/render", handlers.RenderPage)
	http.HandleFunc("/generate-daily", handlers.GenerateDailyData)
	http.HandleFunc("/generate-summary", handlers.GenerateSummaryData)
	http.HandleFunc("/upload-mainpage", handlers.UploadMainPage)
	// http.HandleFunc("/daily", handlers.SlackHandler)
	// http.HandleFunc("/draw", handlers.DrawGraph)
	// http.HandleFunc("/drawstate", handlers.DrawGraphState)

	// http.HandleFunc("/test", handlers.DrawGraphUSMAP)

	// http.HandleFunc("/testrefactor", handlers.TestRefactor)
	fmt.Println("Listening..")
	http.ListenAndServe(":8080", nil)

}
