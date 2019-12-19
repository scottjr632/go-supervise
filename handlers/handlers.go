package handlers

import (
	"context"
	"encoding/json"
	"go-supervise/db"
	"go-supervise/jwt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handlers struct {
	Routable
	repo db.DB
	jwt  *jwt.JWT
	// services services.Services
}

// Handlers ...
type Handlers interface {
	Build() error
}

// NewHandlers returns a new handler
func NewHandlers(server Routable, repo db.DB, j *jwt.JWT) Handlers {
	return &handlers{server, repo, j}
}

func convertToGinHandler(handler func(c context.Context, w http.ResponseWriter, r *http.Request)) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(c, c.Writer, c.Request)
	}
}

func (h *handlers) Build() error {
	api := h.Group("/api")
	protectedRoutes := api.Group("/protected")
	protectedRoutes.Use(func() gin.HandlerFunc {
		return func(c *gin.Context) {
			if err := h.jwt.JWTMiddleware(c.Request, c.Set); err != nil {
				writeError(c.Writer, err, http.StatusUnauthorized)
				c.Abort()
			} else {
				c.Next()
			}
		}
	}())
	workers := api.Group("/workers")
	{
		h.buildWorkerHandlers(workers)
		health := workers.Group("/health")
		{
			h.buildCheckUpHandlers(health)
		}
	}
	auth := api.Group("/auth")
	{
		h.buildAuthHandlers(auth, protectedRoutes)
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
