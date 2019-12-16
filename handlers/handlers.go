package handlers

import (
	"encoding/json"
	"go-supervise/db"
	"io/ioutil"
	"net/http"
)

type handlers struct {
	Routable
	repo db.DB
	// jwt      *jwt.JWT
	// services services.Services
}

type Handlers interface {
	Build() error
}

func NewHandlers(server Routable) Handlers {
	return &handlers{Routable: server}
}

func (h *handlers) Build() error {
	api := h.GROUP("/api")
	{
		buildTestHandlers(api)
	}
	workers := api.GROUP("/workers")
	{
		h.buildWorkerHandlers(workers)
		health := workers.GROUP("/health")
		{
			h.buildCheckUpHandlers(health)
		}
	}
	return nil
}

func writeJSON(w http.ResponseWriter, model interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	jsonModel, err := json.Marshal(model)
	if err != nil {
		return err
	}

	if _, err := w.Write(jsonModel); err != nil {
		return err
	}

	w.WriteHeader(200)
	return nil
}

func readJSON(r *http.Request, model interface{}) error {
	bod, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return err
	}

	err = json.Unmarshal(bod, model)
	return err
}

func writeError(w http.ResponseWriter, err error, code int) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	errJSON := &struct {
		Error struct {
			Message string `json:"message"`
		} `json:"error"`
	}{}

	errJSON.Error.Message = err.Error()
	writeJSON(w, errJSON)
}
