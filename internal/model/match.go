package model

type Match struct {
	TeamA  string `json:"teamA"`
	TeamB  string `json:"teamB"`
	ID     int    `json:"id"`
	Score  Score  `json:"score"`
	Event  string `json:"event"`
	Minute int    `json:"minute"`
	Second int    `json:"second"`
}

type Score struct {
	TeamA   int `json:"teamA-score"`
	TeamB   int `json:"teamB-score"`
	MatchID int `json:"match-id"`
}
