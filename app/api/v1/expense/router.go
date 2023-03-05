package expense

import (
	"github.com/HiDoYa/rent-cat/app/svc/router"
	"github.com/gin-gonic/gin"
)

// Controllable is an interface for implementations of API endpoints
type Controllable interface {
	GetExpense(c *gin.Context)
	GetExpenses(c *gin.Context)
	GetExpenseSummary(c *gin.Context)
}

// Routes adds expenses api endpoints
// Controller injection enables easier testing
func Routes(engine router.Router, controllable Controllable) {
	engine.GET("/expense/:expense_type/:year/:month", controllable.GetExpenses)
	engine.GET("/expenses/:year/:month", controllable.GetExpenses)
	engine.GET("/expense-summary/:year/:month", controllable.GetExpenseSummary)
}