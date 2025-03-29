package models

type Phase struct {
	ID           string           `json:"id"`
	Order        int              `json:"order"`
	Players      []Player         `json:"players"`
	Tables       []Table          `json:"tables"`
	Settings     Specification    `json:"settings"`
	StartTime    *int64           `json:"startTime,omitempty"`
	FinishTime   *int64           `json:"finishTime,omitempty"`
	SkippedTime  *int64           `json:"skippedTime,omitempty"`
	Rounds       []Round          `json:"rounds"`
	FixtureSwaps []FixtureSwap    `json:"fixtureSwaps"`
	EventLog     []PhaseEvent     `json:"eventLog"`
	Ranking      []PlayerRecord   `json:"ranking"`
	TieBreakers  []TieBreakerInfo `json:"tieBreakers"`
}
