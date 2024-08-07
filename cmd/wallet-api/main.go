package main

import (
	"net/http"
	"os"
	"wallet-api/internal/handlers"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	router := handlers.SetupRouter()
	http.ListenAndServe(os.Getenv("HTTP_SERVER_ADDRESS"), router)
}
