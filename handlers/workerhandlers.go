package handlers

import (
	"context"
	"errors"
	"go-supervise/entities"
	"go-supervise/helpers"
	services "go-supervise/services/checkup"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handlers) buildWorkerHandlers(group *gin.RouterGroup) {
	{
		group.POST("/", convertToGinHandler(addWorker))
		group.DELETE("/", convertToGinHandler(removeWorker))
		group.GET("/", convertToGinHandler(getWorkers))
	}
}

func getWorkers(c context.Context, w http.ResponseWriter, r *http.Request) {
	workerID := r.URL.Query().Get("workerId")
	if workerID != "" {
		_, worker := services.GetCheckUpService().FindWorkerByID(workerID)
		if worker != nil {
			writeJSON(w, worker)
		} else {
			writeError(w, services.ErrWorkerNotFound, http.StatusBadRequest)
		}
	} else {
		writeJSON(w, services.GetCheckUpService().GetWorkers())
	}
}

func removeWorker(c context.Context, w http.ResponseWriter, r *http.Request) {
	workerID := r.URL.Query().Get("workerId")
	if workerID != "" {
		if err := services.GetCheckUpService().DeleteWorker(workerID); err != nil {
			writeError(w, err, http.StatusBadRequest)
			return
		}
	} else {
		writeError(w, errors.New("Worker ID not specified"), http.StatusBadRequest)
	}
}

func addWorker(c context.Context, w http.ResponseWriter, r *http.Request) {
	worker := &entities.Worker{}
	if err := readJSON(r, worker); err != nil {
		writeError(w, err, http.StatusBadRequest)
		return
	}

	if err := helpers.AddWorker(worker, services.GetCheckUpService()); err != nil {
		writeError(w, err, http.StatusBadRequest)
	} else {
		writeJSON(w, worker)
	}
}
