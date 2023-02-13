package models

import "time"

// ExpenseType represents an expense type
type ExpenseType struct {
	TypeName string
	CreatedAt time.Time
}

// Expense represents a single expense
type Expense struct {
	ExpenseID int
	ExpenseType string
	Amount float32
	CreatedAt time.Time
}