package handlers

import (
	"github.com/gin-gonic/gin"
)

// Routable represents a serve that has routes
// this is for GIN
type Routable interface {
	Group(relativePath string, handlers ...gin.HandlerFunc) *gin.RouterGroup

	GET(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes
	POST(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes
	PUT(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes
	DELETE(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes
	PATCH(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes
}
