package model

//Response represents the game output to other players
type Response struct {
	Possibilities []Possibility `json:"possibilities"`
}

//Possibility item detected by Robot
type Possibility struct {
	Text     string  `json:"text"`
	Accuracy float64 `json:"accuracy"`
}

//ResponseKnowledge represents knowledge name
type ResponseKnowledge struct {
	Name string `json:"name"`
}
