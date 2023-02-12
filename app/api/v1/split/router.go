package split

import "github.com/gin-gonic/gin"

// Controllable ...
type Controllable interface {
	GetSplit(c *gin.Context)
	PostSplit(c * gin.Context)
}


// Router ...
type Router interface {
	GET(string, ...gin.HandlerFunc) gin.IRoutes
	POST(string, ...gin.HandlerFunc) gin.IRoutes
}

// Routes ...
func Routes(engine Router, controllable Controllable) {
	engine.GET("/split", controllable.GetSplit)
	engine.POST("/split", controllable.PostSplit)
}