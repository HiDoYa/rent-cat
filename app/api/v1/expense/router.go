package expense

import (
	"github.com/HiDoYa/rent-cat/app/svc/router"
	"github.com/gin-gonic/gin"
)

// Controllable is an interface for implementations of API endpoints
type Controllable interface {
	GetExpenses(c *gin.Context)
	GetExpense(c *gin.Context)
}

// Routes adds expenses api endpoints
// Controller injection enables easier testing
func Routes(engine router.Router, controllable Controllable) {
	engine.GET("/expenses", controllable.GetExpenses)
	engine.GET("/expense/:year/:month", controllable.GetExpense)
}