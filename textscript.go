package textscript

import (
	"strings"
)

const ensureValueIsBiggerThanOne = float64(1)

// Run executes a text script that returns a "coefficient" for a word in
// a text based on two parameters: proximity of the word to the beginning
// of the text and distribution of the word across the text.
func Run(text []byte, targetWord string) (result float64) {
	strText := strings.ToLower(string(text))
	targetWordLower := strings.ToLower(targetWord)

	if strText == "" || targetWordLower == "" {
		return
	}

	wordArray := strings.Split(strText, " ")
	var indexArray []int

	for index, word := range wordArray {
		if strings.Contains(word, targetWordLower) {
			indexArray = append(indexArray, index)
		}
	}

	if len(indexArray) == 0 {
		return
	}

	result = ensureValueIsBiggerThanOne

	for i, wordIndex := range indexArray {

		// Proximity index
		result *= ((float64(len(wordArray)) - float64(wordIndex)) / float64(len(wordArray))) +
			ensureValueIsBiggerThanOne

		if i == len(indexArray)-1 {
			break
		}

		// Distribution index
		result *= float64(indexArray[i+1]-wordIndex)/float64(len(wordArray)) +
			ensureValueIsBiggerThanOne
	}

	result -= ensureValueIsBiggerThanOne

	return
}
