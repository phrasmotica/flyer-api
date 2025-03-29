package models

// this, and other models, generated from UI interfaces by:
// https://www.codeconvert.ai/typescript-to-golang-converter

type Flyer struct {
	ID         string         `json:"id"`
	StartTime  *int64         `json:"startTime,omitempty"`
	FinishTime *int64         `json:"finishTime,omitempty"`
	Phases     []Phase        `json:"phases"`
	Ranking    []PlayerRecord `json:"ranking"`
}
