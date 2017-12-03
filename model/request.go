package model

//Request represents the game input
type Request struct {
	Knowledge        Knowledge `json:"knowledge"`
	Text             string    `json:"text"`
	NonExistentChars string    `json:"non_existent_chars"`
	Tip              string    `json:"tip"`
}
