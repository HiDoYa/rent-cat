package expense

import (
	"net/http"
	"strconv"

	"github.com/HiDoYa/rent-cat/app/models"
	"github.com/gin-gonic/gin"
)

// Controller holds implementations of API endpoints
type Controller struct {}

// GetExpenses returns all expenses
// @Summary Get all expenses
// @Schemes
// @Description Get all expenses
// @Tags expenses
// @Accept json
// @Produce json
// @Success 200 {array} models.Expense
// @Router /expenses [get]
func (cont Controller) GetExpenses(c *gin.Context) {
	// TODO Get from database

	var expenses []models.Expense
	for i := 0; i < 5; i++ {
		expense := models.Expense{
			Amount: 0.5,
		}

		expenses = append(expenses, expense)
	}

	c.JSON(http.StatusOK, gin.H{
		"expenses": expenses,
	})
}

// GetExpense gets expenses for a certain month and year
// @Summary gets expenses for a certain month and year
// @Schemes
// @Description gets expenses for a certain month and year
// @Tags expenses
// @Accept json
// @Produce json
// @Success 200 {object} models.Expense
// @Router /expense/:year/:month [get]
func (cont Controller) GetExpense(c *gin.Context) {
	yearStr := c.Param("year")
	monthStr := c.Param("month")

	strconv.Atoi(monthStr)
	strconv.Atoi(yearStr)

	// TODO Get from database

	expense := models.Expense{
		Amount: 5.0,
	}

	c.JSON(http.StatusOK, gin.H{
		"expense": expense,
	})
}