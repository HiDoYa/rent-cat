package api

import (
	"github.com/HiDoYa/rent-cat/app/api/v1/expense"
	"github.com/HiDoYa/rent-cat/app/api/v1/split"
	"github.com/gin-gonic/gin"
)

// Routes adds v1 api endpoints
func Routes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		split.Routes(v1, split.Controller{})
		expense.Routes(v1, expense.Controller{})
	}
}