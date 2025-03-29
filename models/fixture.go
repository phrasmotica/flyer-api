package models

type Fixture struct {
	ID             string          `json:"id"`
	ParentFixtures []ParentFixture `json:"parentFixtures"`
	Scores         []Score         `json:"scores"`
	TableID        string          `json:"tableId"`
	BreakerID      string          `json:"breakerId"`
	StartTime      *int64          `json:"startTime,omitempty"`
	FinishTime     *int64          `json:"finishTime,omitempty"`
	CancelledTime  *int64          `json:"cancelledTime,omitempty"`
	Comment        string          `json:"comment"`
	IsExcluded     bool            `json:"isExcluded"`
}

type ParentFixture struct {
	FixtureID string `json:"fixtureId"`
	TakeLoser bool   `json:"takeLoser"`
}

type Score struct {
	PlayerID string `json:"playerId"`
	Score    int    `json:"score"`
	Runouts  int    `json:"runouts"`
	IsBye    bool   `json:"isBye"`
}
