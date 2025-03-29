package models

type TieBreakerInfo struct {
	ID      string         `json:"id"`
	Name    string         `json:"name"`
	Index   int            `json:"index"`
	ForRank int            `json:"forRank"`
	Records []PlayerRecord `json:"records"`
	Players []Player       `json:"players"`
}
