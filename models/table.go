package models

type Table struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	CostPerHour float64 `json:"costPerHour"`
}
