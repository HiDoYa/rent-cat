package database

import (
	"fmt"

	"github.com/HiDoYa/rent-cat/app/models"
)

// InsertExpense adds a new expense
func (db *Client) InsertExpense(expense models.Expense) (int, error) {
	stmt := "INSERT INTO expense (expense_type, amount) VALUES ($1, $2)"
	result, err := db.client.Exec(stmt, expense.ExpenseType, expense.Amount)

	if err != nil{
		return -1, fmt.Errorf("DB InsertExpense error: %w", err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil{
		return -1, fmt.Errorf("DB InsertExpense error: %w", err)
	}

	return int(lastInsertID), err
}

// SelectExpense retrieves a single expense from the specified month and year
func (db *Client) SelectExpense(expenseType string, month int, year int) (models.Expense, error) {
	stmt := "SELECT expense_id, expense_type, amount, created_at FROM expense " +
		"WHERE EXTRACT('month' from created_at) = $1 " +
		"AND EXTRACT('year' from created_at) = $2" + 
		"AND expense_type = $3"

	rows, err := db.client.Query(stmt, month, year, expenseType)

	if err != nil {
		return models.Expense{}, fmt.Errorf("DB SelectExpense error: %w", err)
	}
	defer rows.Close()

	var expense models.Expense
	err = rows.Scan(&expense.ExpenseID, &expense.ExpenseType, &expense.Amount, &expense.CreatedAt)
	if err != nil {
		return models.Expense{}, fmt.Errorf("DB SelectExpense error: %w", err)
	}

	return expense, nil
}

// SelectExpenses retrieves all expenses from the specified month and year
func (db *Client) SelectExpenses(month int, year int) ([]models.Expense, error) {
	stmt := "SELECT expense_id, expense_type, amount, created_at FROM expense " +
		"WHERE EXTRACT('month' from created_at) = $1 " +
		"AND EXTRACT('year' from created_at) = $2"

	rows, err := db.client.Query(stmt, month, year)

	if err != nil {
		return nil, fmt.Errorf("DB SelectExpenses error: %w", err)
	}
	defer rows.Close()

	allExpenses := []models.Expense{}
	for rows.Next() {
		var currentExpense models.Expense
		err = rows.Scan(&currentExpense.ExpenseID, &currentExpense.ExpenseType, &currentExpense.Amount, &currentExpense.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("DB SelectExpenses error: %w", err)
		}
		allExpenses = append(allExpenses, currentExpense)
	}

	return allExpenses, nil
}

// DeleteExpense removes an expense
func (db *Client) DeleteExpense(expenseID int) {

}
