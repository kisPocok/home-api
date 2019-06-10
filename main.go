package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	println("Home API started.")

	port := os.Getenv("PORT")

	api := NewHomeAPI()
	log.Fatal(http.ListenAndServe(":"+port, api.Router()))
}

// NewHomeAPI creates instance
func NewHomeAPI() HomeAPI {
	return HomeAPI{}
}

type HomeAPI struct{}

// Router returns a mux router
func (api *HomeAPI) Router() *mux.Router {
	r := mux.NewRouter()

	// API Base path
	s := r.PathPrefix("/home/v1/").Subrouter()

	// API Endpoints
	s.HandleFunc("/heartbeat", api.heartBeatReport).Methods(http.MethodPost)

	http.Handle("/", r)
	return r
}

func (api *HomeAPI) heartBeatReport(w http.ResponseWriter, r *http.Request) {

	// TODO save the actual heart beat

	w.WriteHeader(http.StatusNoContent)
	_, _ = fmt.Fprintf(w, "{}")
}
