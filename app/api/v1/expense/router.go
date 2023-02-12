package expense

import "github.com/gin-gonic/gin"

// Controllable ...
type Controllable interface {
	GetExpenses(c *gin.Context)
	GetExpense(c *gin.Context)
}

// Router ...
type Router interface {
	GET(string, ...gin.HandlerFunc) gin.IRoutes
}

// Routes ...
func Routes(engine Router, controllable Controllable) {
	engine.GET("/expenses", controllable.GetExpenses)
	engine.GET("/expense/:year/:month", controllable.GetExpense)
}