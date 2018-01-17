package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func EnterGame(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	buCode := queryValues.Get("buCode")
	id := queryValues.Get("id")
	fmt.Fprintln(w, "buCode:", buCode, " id:", id)
}

func PlayGame(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	buCode := queryValues.Get("buCode")
	id := queryValues.Get("id")
	fmt.Fprintln(w, "buCode:", buCode, " id:", id)
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

	resultEntites := ResultEntities{
		ResultEntity{
			CreditPayout:      0,
			Multiplier:        1,
			Paylines:          2,
			Entity:            "S001",
			TotalCreditPayout: 0,
		},
	}

	gamePlay := GamePlay{
		Entities:          entities,
		ResultEntities:    resultEntites,
		TotalCreditPayout: 0,
	}
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
	body := Body{
		BaseCredit:        50,
		CoinSize:          0.1,
		CurrentMode:       0,
		GameState:         5,
		MaxPayout:         0,
		MeetMaxPayout:     false,
		MemberBalance:     1037053.8,
		MemberCredits:     10370538,
		Modes:             modes,
		Multiplier:        1,
		TotalCreditPayout: 0,
	}
	status := GameStatus{
		Body:     body,
		Messages: "",
		Status:   "Success",
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(status); err != nil {
		panic(err)
	}
}

/*
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}

func EnterGame(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	buCode := queryValues.Get("buCode")
	id := queryValues.Get("id")
	fmt.Fprintln(w, "buCode:", buCode, " id:", id)
}*/
