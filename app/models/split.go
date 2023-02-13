package models

import "time"

// SplitSpecifier are the fields used to specify a new split
type SplitSpecifier struct {
	MyPercentage float32
	HerPercentage float32

	MyNetIncome float32
	HerNetIncome float32
}

// Split represents the expense split between me/her
type Split struct {
	SplitID int
	MyPercentage float32
	HerPercentage float32
	CreatedAt time.Time
}
