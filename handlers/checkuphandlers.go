package handlers

import (
	"context"
	"go-supervise/helpers"
	"go-supervise/services/checkup"
	"net/http"
)

func (h *handlers) buildCheckUpHandlers(group Routable) {
	group.GET("/", convertToGinHandler(h.getWorkersHealth))
}

func (h *handlers) getWorkersHealth(c context.Context, w http.ResponseWriter, r *http.Request) {
	workerID := r.URL.Query().Get("workerId")
	if workerID != "" {
		if health, err := helpers.GetHealthByWorkerID(workerID, h.repo); err != nil {
			writeError(w, err, http.StatusBadRequest)
		} else {
			if err := writeJSON(w, helpers.HealthWrapper{Status: health.Status(), Worker: health.Worker}); err != nil {
				writeError(w, err, http.StatusInternalServerError)
			}
		}
	} else {
		if healthStatus, err := helpers.GetHealthForAllWorkers(checkup.GetCheckUpService(), h.repo); err != nil {
			writeError(w, err, http.StatusInternalServerError)
		} else {
			var wrappedHealths []*helpers.HealthWrapper
			for _, health := range healthStatus {
				wrappedHealth := &helpers.HealthWrapper{Status: health.Status(), Worker: health.Worker, CheckUps: health.Checkups}
				wrappedHealths = append(wrappedHealths, wrappedHealth)
			}

			writeJSON(w, wrappedHealths)
		}
	}
}
