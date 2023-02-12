package models

// Split ...
type Split struct {
	MyPercentage float32
	HerPercentage float32

	Year int
	Month int
}

// SplitSpecifier ...
type SplitSpecifier struct {
	MyPercentage float32
	HerPercentage float32

	MyNetIncome float32
	HerNetIncome float32
}