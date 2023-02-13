package database

import (
	"fmt"

	"github.com/HiDoYa/rent-cat/app/models"
)

// InsertExpenseType adds a new expense type
func (db *Client) InsertExpenseType(expenseType models.ExpenseType) (int, error) {
	stmt := "INSERT INTO expense_type (expense_type) VALUES ($1)"
	result, err := db.client.Exec(stmt, expenseType.TypeName)
	if err != nil{
		return 0, fmt.Errorf("DB InsertExpenseType error: %w", err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil{
		return -1, fmt.Errorf("DB InsertExpense error: %w", err)
	}

	return int(lastInsertID), err
}

// SelectExpenseTypes retrieves all active expense types
func (db *Client) SelectExpenseTypes() ([]models.ExpenseType, error) {
	stmt := "SELECT type_name, created_at FROM expense_type " +
		"WHERE active = true"

	rows, err := db.client.Query(stmt)

	if err != nil {
		return nil, fmt.Errorf("DB SelectExpenseTypes error: %w", err)
	}
	defer rows.Close()

	allExpenseTypes := []models.ExpenseType{}
	for rows.Next() {
		var currentExpenseType models.ExpenseType
		err = rows.Scan(&currentExpenseType.TypeName, &currentExpenseType.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("DB SelectExpenseTypes error: %w", err)
		}
		allExpenseTypes = append(allExpenseTypes, currentExpenseType)
	}

	return allExpenseTypes, nil
}