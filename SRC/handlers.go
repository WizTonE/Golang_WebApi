package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello!")
}

func IndexQuery(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintln(w, "Hello! ", id)
}

func EnterGame(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	buCode := queryValues.Get("buCode")
	id := queryValues.Get("id")
	fmt.Fprintln(w, "buCode:", buCode, " id:", id)
}

func PlayGame(w http.ResponseWriter, r *http.Request) {
	gamePlay := GamePlay{
		TotalCreditPayout: 0,
	}
	gamePlay.Entities = gamePlay.GetEntities()
	gamePlay.ResultEntities = gamePlay.GetResultEntities()

	body := Body{
		BaseCredit:        50,
		CoinSize:          0.1,
		CurrentMode:       0,
		GameState:         5,
		MaxPayout:         0,
		MeetMaxPayout:     false,
		MemberBalance:     1037053.8,
		MemberCredits:     10370538,
		Multiplier:        1,
		TotalCreditPayout: 0,
	}
	body.Modes = body.GetModesStatus(gamePlay)

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
