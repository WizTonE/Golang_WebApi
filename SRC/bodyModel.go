package main

type Body struct {
	BaseCredit        int     `json :"BaseCredit"`
	CoinSize          float32 `json:"CoinSize"`
	CurrentMode       int     `json:CurrentMode`
	GameState         int     `json:GameState`
	MaxPayout         int     `json:MaxPayout`
	MeetMaxPayout     bool    `json:MeetMaxPayout`
	MemberBalance     float64 `json:MemberBalance`
	MemberCredits     int64   `json:MemberCredits`
	Modes             Modes   `json:Modes`
	Multiplier        int     `json:Multiplier`
	TotalCreditPayout int     `json:TotalCreditPayout`
}
