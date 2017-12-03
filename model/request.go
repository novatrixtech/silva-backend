package model

//Request represents the game input
type Request struct {
	Matter           Matter `json:"matter"`
	Text             string `json:"text"`
	NonExistentChars string `json:"non_existent_chars"`
	Tip              string `json:"tip"`
}
