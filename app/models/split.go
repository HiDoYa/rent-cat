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

// SplitType represents how an expense should be split
type SplitType int

const (
	// FiftyFiftySplit splits the expense 50/50
	FiftyFiftySplit SplitType = iota
	// DefaultSplit uses the configured split in the database
	DefaultSplit
	// HerOnlySplit only makes her pay
	HerOnlySplit
	// MeOnlySplit only makes me pay
	MeOnlySplit
)

func (s SplitType) computeSplit(defaultSplit Split) Split {
	switch s {
	case FiftyFiftySplit:
		return Split{
			MyPercentage: 50,
			HerPercentage: 50,
		}
	case DefaultSplit:
		return defaultSplit
	case HerOnlySplit:
		return Split{
			HerPercentage: 100,
		}
	case MeOnlySplit:
		return Split{
			MyPercentage: 100,
		}
	}

	return defaultSplit
}