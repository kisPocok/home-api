package main

import (
	"net/http"
	"os"

	"github.com/kisPocok/home-api/home"
	"github.com/kisPocok/home-api/logger"
	"go.uber.org/zap"
)

func main() {
	println("Home API started.")

	port := os.Getenv("PORT")

	log := logger.New()
	log.Info("Home API started", zap.String("port", port))

	api := home.NewHomeAPI(log)
	err := http.ListenAndServe(":"+port, api.Router())
	if err != nil {
		log.Fatal("Could not start HTTP server", zap.Error(err))
	}
}
