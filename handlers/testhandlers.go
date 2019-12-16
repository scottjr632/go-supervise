package handlers

import (
	"context"
	"net/http"
)

func buildTestHandlers(base Routable) {
	base.GET("/test", sendTest)
}

func sendTest(c context.Context, w http.ResponseWriter, r *http.Request) {
	writeJSON(w, &struct {
		Message string `json:"message"`
	}{
		"Hello, there man!",
	})
}
