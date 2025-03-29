package models

type PlayerRecord struct {
	PlayerID   string `json:"playerId"`
	Name       string `json:"name"`
	Played     int    `json:"played"`
	Wins       int    `json:"wins"`
	Draws      int    `json:"draws"`
	Losses     int    `json:"losses"`
	Diff       int    `json:"diff"`
	Runouts    int    `json:"runouts"`
	Points     int    `json:"points"`
	Incomplete bool   `json:"incomplete"`
	Rank       int    `json:"rank"`
	TieBroken  bool   `json:"tieBroken"`
}
