package models

type PhaseEvent struct {
	Level     PhaseEventLevel `json:"level"`
	Message   string          `json:"message"`
	Timestamp int64           `json:"timestamp"`
}

type PhaseEventLevel int

const (
	Default  PhaseEventLevel = iota
	Internal                 = 100
)
