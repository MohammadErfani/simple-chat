package main

import (
	"chat/internal/handlers"
	"log"
	"net/http"
)

func main() {
	routes := routes()
	log.Println("Starting channel listener")
	go handlers.ListenToWsChannel()
	log.Println("Start web server on port 8080")
	_ = http.ListenAndServe(":8080", routes)
}
