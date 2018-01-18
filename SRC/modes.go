package main

type Mode struct {
	CurrentPlay  int      `json:CurrentPlay`
	CurrentPoint int      `json:CurrentPoint`
	CurrentRound int      `json:CurrentRound`
	GamePlay     GamePlay `json:GamePlay, omitempty`
	TotalPlay    int      `json:TotalPlay`
	TotalPoint   int      `json:TotalPoint`
	TotalRound   int      `json:TotalRound`
	Type         int      `json:Type`
}

type Modes []Mode
