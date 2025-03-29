package models

type Specification struct {
	MatchLengthModel       MatchLengthModel `json:"matchLengthModel"`
	BestOf                 int              `json:"bestOf"`
	RaceTo                 int              `json:"raceTo"`
	WinsRequired           int              `json:"winsRequired"`
	RuleSet                RuleSet          `json:"ruleSet"`
	Format                 Format           `json:"format"`
	StageCount             int              `json:"stageCount"`
	RandomlyDrawAllRounds  bool             `json:"randomlyDrawAllRounds"`
	RequireCompletedRounds bool             `json:"requireCompletedRounds"`
	AllowEarlyFinish       bool             `json:"allowEarlyFinish"`
	TieBreaker             TieBreaker       `json:"tieBreaker"`
	EntryFeeRequired       bool             `json:"entryFeeRequired"`
	EntryFee               float64          `json:"entryFee"`
	MoneySplit             MoneySplit       `json:"moneySplit"`
	Name                   string           `json:"name"`
}

type MatchLengthModel int

const (
	Fixed MatchLengthModel = iota
	Variable
)

type Format int

const (
	Knockout Format = iota
	RoundRobin
	WinnerStaysOn
)

type RuleSet int

const (
	Blackball RuleSet = iota
	International
)

type TieBreaker int

const (
	HeadToHead TieBreaker = iota
	Runouts
)

type MoneySplit int

const (
	WinnerTakesAll MoneySplit = iota
	SeventyThirty
	SixtyTwentyFiveFifteen
	SemiFinalists
)
