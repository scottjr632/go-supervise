package server

import (
	"fmt"
	"go-supervise/handlers"

	"github.com/gin-gonic/gin"
)

// NewServer returns a new server instance
func NewServer(config *Config) Server {
	applyDefaults(config)
	return &server{Config: config}
}

type server struct {
	*Config
	*group
	E *gin.Engine
}

type group struct {
	// gin *gin.Engine
	gin *gin.RouterGroup
}

func convertToGinHandler(handlers ...handlers.ServerHandler) []gin.HandlerFunc {
	ginHandlers := make([]gin.HandlerFunc, 0)
	for _, handler := range handlers {
		ginHandlers = append(ginHandlers, func(c *gin.Context) {
			handler(c, c.Writer, c.Request)
		})
	}
	return ginHandlers
}

func (s *group) POST(path string, handlers ...handlers.ServerHandler) {
	s.gin.POST(path, convertToGinHandler(handlers...)...)
}

func (s *group) GET(path string, handlers ...handlers.ServerHandler) {
	s.gin.GET(path, convertToGinHandler(handlers...)...)
}

func (s *group) DELETE(path string, handlers ...handlers.ServerHandler) {
	s.gin.DELETE(path, convertToGinHandler(handlers...)...)
}

func (s *group) PATCH(path string, handlers ...handlers.ServerHandler) {
	s.gin.DELETE(path, convertToGinHandler(handlers...)...)
}

func (s *group) GROUP(path string, handlers ...handlers.ServerHandler) handlers.Routable {
	g := s.gin.Group(path, convertToGinHandler(handlers...)...)
	return &group{g}
}

func (s *server) Build() Server {
	g := gin.Default()

	s.E = g
	s.group = &group{gin: &g.RouterGroup}

	return s
}

// BuildAndrun ...
func (s *server) Run() error {
	return s.E.Run(fmt.Sprintf(":%v", s.Port))
}
