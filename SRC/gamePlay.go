package main

type GamePlay struct {
	Entities          Entities       `json:Entities`
	ResultEntities    ResultEntities `json:ResultEntities`
	TotalCreditPayout int            `json:TotalCreditPayout`
}
