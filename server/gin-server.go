package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/scottjr632/go-bloom-filter"
)

// NewServer returns a new server instance
func NewServer(config *Config) Server {
	applyDefaults(config)
	return &server{Config: config}
}

type server struct {
	*Config
	*gin.Engine
}

func CORS() gin.HandlerFunc {
	bloomfilter.
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:1234")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, access-control-allow-origin")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func (s *server) Build() Server {
	g := gin.Default()

	g.Use(CORS())

	s.Engine = g

	return s
}

// BuildAndrun ...
func (s *server) Start() error {
	return s.Run(fmt.Sprintf(":%v", s.Port))
}
