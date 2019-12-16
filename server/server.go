package server

import (
	"go-supervise/handlers"
)

type Buildable interface {
	Build() Server
	Run() error
}

// Server ...
type Server interface {
	Buildable
	handlers.Routable
}
