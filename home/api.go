package home

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// NewHomeAPI creates instance
func NewHomeAPI() Api {
	return Api{}
}

type Api struct{}

// Router returns a mux router
func (api *Api) Router() *mux.Router {
	r := mux.NewRouter()

	// Api Base path
	s := r.PathPrefix("/home/v1/").Subrouter()

	// Api Endpoints
	s.HandleFunc("/heartbeat", api.heartBeatReport).Methods(http.MethodPost)

	http.Handle("/", r)
	return r
}

func (api *Api) heartBeatReport(w http.ResponseWriter, r *http.Request) {

	// TODO save the actual heart beat

	w.WriteHeader(http.StatusNoContent)
	_, _ = fmt.Fprintf(w, "{}")
}
