package expense

import (
	"net/http"
	"strconv"

	"github.com/HiDoYa/rent-cat/app/models"
	"github.com/HiDoYa/rent-cat/app/svc/database"
	"github.com/gin-gonic/gin"
)

// Controller holds implementations of API endpoints
type Controller struct {}

// GetExpenseSummary gets expense summary for a certain month and year
// @Summary gets expense summary for a certain month and year
// @Schemes
// @Description gets expense summary for a certain month and year
// @Tags expenses
// @Accept json
// @Produce json
// @Param year path string true "current year"
// @Param month path string true "current month"
// @Success 200 {object} models.ExpenseSummary
// @Router /expense-summary/{year}/{month} [get]
func (cont Controller) GetExpenseSummary(c *gin.Context) {
	yearStr := c.Param("year")
	monthStr := c.Param("month")

	month, err := strconv.Atoi(monthStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	client, err := database.MakeDb()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	expenses, err := client.SelectExpenses(month, year)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	split, err := client.SelectLatestSplit()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	expenseSummary := models.CreateExpenseSummary(expenses, split)

	c.JSON(http.StatusOK, gin.H{
		"summary": expenseSummary,
	})
}

// GetExpense gets expenses for a certain month and year
// @Summary gets expenses for a certain month and year
// @Schemes
// @Description gets expenses for a certain month and year
// @Tags expenses
// @Accept json
// @Produce json
// @Param expense_type path string true "expense type"
// @Param year path string true "current year"
// @Param month path string true "current month"
// @Success 200 {object} models.Expense
// @Router /expense/{expense_type}/{year}/{month} [get]
func (cont Controller) GetExpense(c *gin.Context) {
	expenseType := c.Param("expenseType")
	yearStr := c.Param("year")
	monthStr := c.Param("month")

	month, err := strconv.Atoi(monthStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	client, err := database.MakeDb()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	expenses, err := client.SelectExpense(expenseType, month, year)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"expenses": expenses,
	})
}

// GetExpenses gets expenses for a certain month and year
// @Summary gets expenses for a certain month and year
// @Schemes
// @Description gets expenses for a certain month and year
// @Tags expenses
// @Accept json
// @Produce json
// @Param year path string true "current year"
// @Param month path string true "current month"
// @Success 200 {array} models.Expense
// @Router /expenses/{year}/{month} [get]
func (cont Controller) GetExpenses(c *gin.Context) {
	yearStr := c.Param("year")
	monthStr := c.Param("month")

	month, err := strconv.Atoi(monthStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	client, err := database.MakeDb()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	expenses, err := client.SelectExpenses(month, year)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"expenses": expenses,
	})
}