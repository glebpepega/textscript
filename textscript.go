package textscript

import (
	"strings"
)

// Run executes a text script that return a "coefficient" for a word in a
// text based on three parameters: proximity of the word to the beginning of the
// text, distribution of the word across the text and frequency of the word in the text.
func Run(text string, targetWord string) (result float64) {
	if text == "" || targetWord == "" {
		return
	}

	wordArray := strings.Split(text, " ")

	var indexArray []int

	for index, word := range wordArray {
		if word == targetWord {
			indexArray = append(indexArray, index)
		}
	}

	if len(indexArray) == 0 {
		return
	}

	result = 1

	const ensureValueIsMoreThanOne = float64(1)

	for i, v := range indexArray {

		// proximity index
		result *= ((float64(len(wordArray)) - float64(v)) / float64(len(wordArray))) +
			ensureValueIsMoreThanOne

		if i == len(indexArray)-1 {
			break
		}

		// distribution index
		result *= float64(indexArray[i+1]-v)/float64(len(wordArray)) +
			ensureValueIsMoreThanOne
	}

	// frequency index
	result *= float64(len(indexArray))/float64(len(wordArray)) +
		ensureValueIsMoreThanOne

	result -= 1

	return
}
