package models

type Round struct {
	Index       int       `json:"index"`
	Name        string    `json:"name"`
	RaceTo      *int      `json:"raceTo,omitempty"`
	IsGenerated bool      `json:"isGenerated"`
	Fixtures    []Fixture `json:"fixtures"`
}
