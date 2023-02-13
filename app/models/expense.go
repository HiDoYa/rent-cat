package models

import "time"

// ExpenseType ...
type ExpenseType struct {
	TypeName string
	CreatedAt time.Time
}

// Expense ...
type Expense struct {
	ExpenseID int
	ExpenseType string
	Amount float32
	CreatedAt time.Time
}