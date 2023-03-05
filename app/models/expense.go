package models

import (
	"time"
)

// ExpenseType represents an expense type
type ExpenseType struct {
	TypeName string
	CreatedAt time.Time
	SplitType SplitType
}

// ExpenseStatus represents the payment status of the expense
type ExpenseStatus string

const (
	// PaidStatus means the expense has been paid by both parties
	PaidStatus ExpenseStatus = "Paid"
	// CoveredStatus means the expense is covered and does not need to be paid
	CoveredStatus ExpenseStatus = "Covered"
	// UnpaidStatus means the expense has not been paid by both parties
	UnpaidStatus ExpenseStatus = "Unpaid"
)

func (es ExpenseStatus) priority() int {
	switch es {
	case PaidStatus:
		return 0
	case CoveredStatus:
		return 1
	case UnpaidStatus:
		return 2
	}

	return -1
}

// Expense represents a single expense
type Expense struct {
	ExpenseID int
	ExpenseType ExpenseType
	Amount float32
	Status ExpenseStatus
	CreatedAt time.Time
}

// ExpenseSummary represents a group of expenses for a month
type ExpenseSummary struct {
	MyAmount float32
	HerAmount float32
	Status string
}

// CreateExpenseSummary ...
func CreateExpenseSummary(expenses []Expense, defaultSplit Split) ExpenseSummary {
	expenseSummary := ExpenseSummary{}

	currentStatus := PaidStatus

	for _, expense := range expenses {
		expenseType := expense.ExpenseType
		split := expenseType.SplitType.computeSplit(defaultSplit)

		expenseSummary.MyAmount += split.MyPercentage / 100.0 * expense.Amount
		expenseSummary.HerAmount += split.HerPercentage / 100.0 * expense.Amount
		if expense.Status.priority() > currentStatus.priority() {
			currentStatus = expense.Status
		}
	}

	expenseSummary.Status = string(currentStatus)
	return expenseSummary
}