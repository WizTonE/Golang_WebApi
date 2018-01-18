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

func (status *Body) GetModesStatus(gamePlay GamePlay) Modes {
	modes := Modes{
		Mode{
			CurrentPlay:  0,
			CurrentPoint: 0,
			CurrentRound: 0,
			GamePlay:     gamePlay,
			TotalPlay:    0,
			TotalPoint:   0,
			TotalRound:   0,
			Type:         0,
		},
		Mode{
			CurrentPlay:  0,
			CurrentPoint: 0,
			CurrentRound: 0,
			TotalPlay:    0,
			TotalPoint:   0,
			TotalRound:   0,
			Type:         1,
		},
	}

	return modes
}
