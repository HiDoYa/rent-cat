package models

import "time"

// SplitSpecifier ...
type SplitSpecifier struct {
	MyPercentage float32
	HerPercentage float32

	MyNetIncome float32
	HerNetIncome float32
}

// Split ...
type Split struct {
	SplitID int
	MyPercentage float32
	HerPercentage float32
	CreatedAt time.Time
}
