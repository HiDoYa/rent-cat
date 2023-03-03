package models

import "time"

// ExpenseType represents an expense type
type ExpenseType struct {
	TypeName string
	CreatedAt time.Time
	SplitType SplitType
}

// Expense represents a single expense
type Expense struct {
	ExpenseID int
	ExpenseType ExpenseType
	Amount float32
	Status int
	CreatedAt time.Time
}

// ExpenseSummary represents a group of expenses for a month
type ExpenseSummary struct {
	MyAmount float32
	HerAmount float32
	Status bool
}

func createExpenseSummary(expenses []Expense, defaultSplit Split) ExpenseSummary {
	expenseSummary := ExpenseSummary{}

	for _, expense := range expenses {
		expenseType := expense.ExpenseType
		split := expenseType.SplitType.computeSplit(defaultSplit)

		expenseSummary.MyAmount += split.MyPercentage / 100.0 * expense.Amount
		expenseSummary.HerAmount += split.HerPercentage / 100.0 * expense.Amount
	}

	return expenseSummary
}