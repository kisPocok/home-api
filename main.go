package main

import (
	"net/http"
	"os"

	"github.com/kisPocok/home-api/home"
	"github.com/kisPocok/home-api/logger"
	"go.uber.org/zap"
)

func main() {
	port := os.Getenv("PORT")
	user := os.Getenv("API_USER")
	pass := os.Getenv("API_PASS")

	log := logger.New()
	log.Info("Home API started", zap.String("port", port))

	api := home.NewHomeAPI(user, pass, log)
	err := http.ListenAndServe(":"+port, api.Router())
	if err != nil {
		log.Fatal("Could not start HTTP server", zap.Error(err))
	}
}
