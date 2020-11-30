package app

import (
	"GoAuth2/main/router"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func StartApplication() {
	loadConfigurationFile()
	log.Printf("server started on port %s", os.Getenv("PORT"))
	listenAndServe()
}

func listenAndServe() {
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), router.CreateRoutes()); err != nil {
		log.Fatalf("%s", err)
	}
}

func loadConfigurationFile() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("failed to load configuration file")
	}
}
