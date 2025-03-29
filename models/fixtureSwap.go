package models

type FixtureSwap struct {
	ID            string `json:"id"`
	RoundAIndex   int    `json:"roundAIndex"`
	FixtureAIndex int    `json:"fixtureAIndex"`
	FixtureAID    string `json:"fixtureAId"`
	RoundBIndex   int    `json:"roundBIndex"`
	FixtureBIndex int    `json:"fixtureBIndex"`
	FixtureBID    string `json:"fixtureBId"`
	Timestamp     int64  `json:"timestamp"`
}
