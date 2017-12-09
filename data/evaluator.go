package data

import (
	"log"
	"strings"

	"github.com/novatrixtech/silva-backend/model"
)

const confidenceFactor float64 = 0.70

//EvaluateWord does all word analysis and classification
func EvaluateWord(text string, matter string) (word model.Word) {
	text = strings.ToLower(strings.Replace(text, " ", "", -1))
	word.Length = len(text)
	word.Text = text
	word.Matter.Name = matter
	return
}

//EvaluateCondition does analysis
func EvaluateCondition(words *[]model.Word, evaluationChars string, nonExistentChars string) (possibilities []model.Possibility) {
	var allOptions, allFoundChars, successChars string
	var predictionAccuracy float64
	log.Println(" ")
	log.Println("===> Evalutating this request: ", evaluationChars, " - NonExistentChars: ", nonExistentChars)
	for _, word := range *words {
		points, otherOptions, foundChars := calcWordValues(word, evaluationChars, nonExistentChars, allFoundChars)
		//IMPORTANT: PAY ATTENTION TO CONFIDENCE FACTOR
		if points >= confidenceFactor {
			poss := model.Possibility{}
			poss.Text = word.Text
			poss.Accuracy = points
			possibilities = append(possibilities, poss)
		}
		if points > 0 {
			log.Println("Word: ", word.Text, " points: ", points, " otherOptions: ", otherOptions, " foundChars: ", foundChars)
			allOptions += otherOptions
			if len(foundChars) > 0 {
				allFoundChars += foundChars
			}
		}
	}
	for _, char := range allFoundChars {
		allOptions = strings.Replace(allOptions, string(char), "", -1)
	}
	successChars, numberOccurrencies := wordCount(allOptions)
	if len(allOptions) > 0 && numberOccurrencies > 0 {
		predictionAccuracy = float64(numberOccurrencies) / float64(len(allOptions))
	}
	log.Println("[EvaluateCondition] Letter: ", successChars, " occurrency: ", numberOccurrencies, " AllOptions: ", allOptions, " len(allOptions): ", len(allOptions), " accuracy: ", predictionAccuracy)
	poss := model.Possibility{}
	poss.Text = successChars
	poss.Accuracy = predictionAccuracy
	possibilities = append(possibilities, poss)

	return
}

func calcWordValues(word model.Word, evaluationChars string, nonExistentChars string, ignoreChars string) (points float64, otherPossibilities string, foundChars string) {
	//If word length and evaluationLetters length doesn't match his value is automatically zero
	if len(word.Text) != len(evaluationChars) {
		return
	}
	for _, charNC := range []byte(nonExistentChars) {
		if strings.Contains(word.Text, string(charNC)) {
			return
		}
	}
	//Starts with 1 because length is a column for statistics purposes
	var numItems = 1 + len(evaluationChars)
	//log.Println("numItems: ", numItems, " word: ", word.Text, " evaluationLetters: ", evaluationChars, " nonExistentChars: ", nonExistentChars, " ignoreChars: ", ignoreChars)
	var columns = make([]int, numItems)
	var indexColumn int
	columns[indexColumn] = 1
	indexColumn++
	for index, char := range []byte(evaluationChars) {
		if word.Text[index] == char {
			columns[indexColumn] = 1
			foundChars += string(char)
		} else {
			otherPossibilities += string(word.Text[indexColumn-1])
		}
		indexColumn++
	}
	var total int
	for _, column := range columns {
		total += column
	}
	for _, char := range foundChars {
		otherPossibilities = strings.Replace(otherPossibilities, string(char), "", -1)
	}
	for _, char := range ignoreChars {
		otherPossibilities = strings.Replace(otherPossibilities, string(char), "", -1)
	}
	//log.Println("foundChars: ", foundChars, " otherPossibilities: ", otherPossibilities, " len(otherPossibilities): ", len(otherPossibilities), " len(evaluationLetters): ", len(evaluationChars))
	points = float64(total) / float64(len(columns))
	return
}

func wordCount(text string) (char string, numberOccurrencies int) {
	var counter int
	var selectedLetter byte
	letters := []byte(strings.Replace(text, " ", "", -1))
	for _, letter := range letters {
		if selectedLetter == 0 {
			selectedLetter = letter
			numberOccurrencies = 0
		}
		counter = 0
		for _, tmp := range letters {
			if letter == tmp {
				counter++
			}
		}
		if counter > numberOccurrencies {
			selectedLetter = letter
			numberOccurrencies = counter
		}
	}
	//log.Println("word: ", text, " selected: ", string(selectedLetter), " occurrencies: ", numberOccurrencies)
	char = string(selectedLetter)
	return
}
