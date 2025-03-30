package contracts

import "phrasmotica/flyer-api/models"

type CreateFlyerRequest struct {
	StartTime  *int64                `json:"startTime,omitempty"`
	FinishTime *int64                `json:"finishTime,omitempty"`
	Phases     []models.Phase        `json:"phases"`
	Ranking    []models.PlayerRecord `json:"ranking"`
}
