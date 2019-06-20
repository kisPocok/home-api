package home

import (
	"fmt"
	"net/http"

	"github.com/goji/httpauth"
	"github.com/gorilla/mux"
	"github.com/kisPocok/home-api/converter"
	"go.uber.org/zap"
)

// NewHomeAPI creates instance
func NewHomeAPI(user, pass string, logger *zap.Logger) Api {
	return Api{
		user: user,
		pass: pass,
		log:  logger,
	}
}

type Api struct {
	user string
	pass string
	log  *zap.Logger
}

// Router returns a mux router
func (api *Api) Router() *mux.Router {
	r := mux.NewRouter()
	authOpts := httpauth.AuthOptions{
		Realm:               "Home-API",
		User:                api.user,
		Password:            api.pass,
		UnauthorizedHandler: api.unauthorizedHandler(),
	}
	r.Use(httpauth.BasicAuth(authOpts)) // Force to use Basic Auth for every request

	// Api Base path
	s := r.PathPrefix("/home/v1/").Subrouter()

	// Api Endpoints
	s.HandleFunc("/heartbeat", api.heartBeatReport).Methods(http.MethodPost)
	s.HandleFunc("/flowers", api.flowerReport).Methods(http.MethodPost)

	http.Handle("/", r)
	return r
}

func (api *Api) heartBeatReport(w http.ResponseWriter, r *http.Request) {
	// Log the usage
	api.log.Info("heart-beat", zap.Any("label", "heartbeat"))

	w.WriteHeader(http.StatusNoContent)
}

func (api *Api) flowerReport(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		api.log.Error("flower internal error")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, "{}")
		return
	}

	f := r.Form
	if f.Get("device") == "" {
		api.log.Error("flower missing device")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintf(w, `{"error":"missing device name"}`)
		return
	}

	api.log.Info("flower-report",
		zap.String("device", f.Get("device")),
		zap.String("mac", f.Get("mac")),
		zap.String("framework", f.Get("framework")),
		zap.Float64("temperature", converter.StrToFloat(f.Get("temperature"))),
		zap.Int("moisture", converter.StrToInt(f.Get("moisture"))),
		zap.Int("light", converter.StrToInt(f.Get("light"))),
		zap.Int("conductivity", converter.StrToInt(f.Get("conductivity"))),
		zap.Int("battery", converter.StrToInt(f.Get("battery"))),
	)

	w.WriteHeader(http.StatusNoContent)
}

func (api *Api) unauthorizedHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		api.log.Error("unauthorized request",
			zap.Any("credentials", r.Header.Get("Authorization")),
		)
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = fmt.Fprintf(w, "{}")
	}
}
