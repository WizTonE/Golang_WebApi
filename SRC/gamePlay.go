package main

type GamePlay struct {
	Entities          Entities       `json:Entities, omitempty`
	ResultEntities    ResultEntities `json:ResultEntities, omitempty`
	TotalCreditPayout int            `json:TotalCreditPayout, omitempty`
}

func (status *GamePlay) GetEntities() Entities {
	entities := Entities{
		Entity{
			Row1: "N002",
			Row2: "N001",
			Row3: "N009",
			Row4: "N007",
		},
		Entity{
			Row1: "N001",
			Row2: "N006",
			Row3: "N009",
			Row4: "N007",
		},
		Entity{
			Row1: "S001",
			Row2: "N002",
			Row3: "S001",
			Row4: "N010",
		},
		Entity{
			Row1: "N005",
			Row2: "N009",
			Row3: "N003",
			Row4: "N004",
		},
	}

	return entities
}

func (status *GamePlay) GetResultEntities() ResultEntities {
	resultEntites := ResultEntities{
		ResultEntity{
			CreditPayout:      0,
			Multiplier:        1,
			Paylines:          2,
			Entity:            "S001",
			TotalCreditPayout: 0,
		},
	}

	return resultEntites
}
