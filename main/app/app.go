package app

import (
	"GoAuth2/main/controllers"
	"log"
	"net/http"
)

func StartApplication() {
	http.HandleFunc("/users", controllers.GetUserByUsername)

	log.Printf("Server starting on port %s", "8080")

	listenAndServe()
}

func listenAndServe() {
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("%s", err)
	}
}
