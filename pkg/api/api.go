package api

import (
	"censor/pkg/censor"
	"censor/pkg/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Программный интерфейс сервера
type API struct {
	censor censor.Service
	router *mux.Router
}

// Конструктор объекта API
func New(censor censor.Service) *API {
	api := API{censor: censor}
	api.router = mux.NewRouter()
	api.endpoints()
	return &api
}

// Маршрутизатор запросов.
func (api *API) Router() *mux.Router {
	return api.router
}

func (api *API) endpoints() {
	api.router.Methods(http.MethodPost).Path("/comments/validate").HandlerFunc(api.validate)
}

func (api *API) validate(w http.ResponseWriter, r *http.Request) {
	var req models.Comment
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Body decoding error", http.StatusBadRequest)
		return
	}

	if !api.censor.Validate(req.Content) {
		http.Error(w, "Bad words were found", http.StatusBadRequest)
	}
}
