package router

import "github.com/gin-gonic/gin"

// Router is an interface representing any gin router that can handle API requests
// e.g. gin.Engine, gin.RouterGroup
type Router interface {
	GET(string, ...gin.HandlerFunc) gin.IRoutes
	POST(string, ...gin.HandlerFunc) gin.IRoutes
	DELETE(string, ...gin.HandlerFunc) gin.IRoutes
}

