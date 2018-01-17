package main

type ResultEntity struct {
	CreditPayout      int    `json:CreditPayout`
	Multiplier        int    `json:Multiplier`
	Paylines          int    `json:Paylines`
	Entity            string `json:Entity`
	TotalCreditPayout int    `json:TotalCreditPayout`
}

type ResultEntities []ResultEntity
