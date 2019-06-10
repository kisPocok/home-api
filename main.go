package main

import (
	"log"
	"net/http"
	"os"

	"github.com/kisPocok/home-api/home"
	"go.uber.org/zap"
)

func main() {
	println("Home API started.")

	port := os.Getenv("PORT")

	logger, _ := zap.NewProduction()
	logger.Info("Home API started", zap.String("port", port))

	api := home.NewHomeAPI(logger)
	log.Fatal(http.ListenAndServe(":"+port, api.Router()))
}
