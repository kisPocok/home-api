package home

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/common/log"
	"go.uber.org/zap"
)

// NewHomeAPI creates instance
func NewHomeAPI(logger *zap.Logger) Api {
	return Api{
		log: logger,
	}
}

type Api struct {
	log *zap.Logger
}

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

	// Log the usage
	log.Info("heart-beat", zap.String("label", "heartbeat"))

	w.WriteHeader(http.StatusNoContent)
	_, _ = fmt.Fprintf(w, "{}")
}
