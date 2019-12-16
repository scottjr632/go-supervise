package handlers

import (
	"context"
	"net/http"
)

type ServerHandler func(c context.Context, w http.ResponseWriter, r *http.Request)

type Routable interface {
	POST(path string, handlers ...ServerHandler)
	GET(path string, handlers ...ServerHandler)
	DELETE(path string, handlers ...ServerHandler)
	PATCH(path string, handlers ...ServerHandler)

	GROUP(path string, handlers ...ServerHandler) Routable
}
