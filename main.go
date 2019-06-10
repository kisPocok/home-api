package main

import (
	"log"
	"net/http"
	"os"

	"github.com/kisPocok/home-api/home"
)

func main() {
	println("Home API started.")

	port := os.Getenv("PORT")

	api := home.NewHomeAPI()
	log.Fatal(http.ListenAndServe(":"+port, api.Router()))
}
