package model

//Word represents the word to be matched at game
type Word struct {
	Text   string `json:"text"`
	Matter Matter `json:"matter"`
	Length int    `json:"length"`
	Facts  []Fact `json:"facts"`
}

//Matter represents the category of a matter that a word belongs to
type Matter struct {
	Name string `json:"name"`
}

//Fact represents facts and general info related a word
type Fact struct {
	Info string `json:"info"`
}
