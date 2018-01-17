package main

type GameStatus struct {
	Body     Body   `json:Body`
	Messages string `json:Messages`
	Status   string `json:Status`
}
