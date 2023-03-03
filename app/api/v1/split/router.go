package split

import (
	"github.com/HiDoYa/rent-cat/app/svc/router"
	"github.com/gin-gonic/gin"
)

// Controllable is an interface for implementations of API endpoints
type Controllable interface {
	GetLatestSplit(c *gin.Context)
	PostSplit(c * gin.Context)
}

// Routes adds split api endpoints
// Controller injection enables easier testing
func Routes(engine router.Router, controllable Controllable) {
	engine.GET("/split", controllable.GetLatestSplit)
	engine.POST("/split", controllable.PostSplit)
}