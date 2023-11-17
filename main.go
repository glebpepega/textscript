package textscript

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("text.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := strings.ToLower(string(content))
	targetWord := strings.ToLower("sample")

	fmt.Println(Run(text, targetWord))
}

func Run(text string, targetWord string) (result float64) {
	result = 1

	wordsArray := strings.Split(text, " ")

	var indexArray []int

	for index, word := range wordsArray {
		if word == targetWord {
			indexArray = append(indexArray, index)
		}
	}

	if len(indexArray) == 0 {
		result = 0

		return
	}

	// ибо нехуй
	const ensureValueIsMoreThanOne = float64(1)

	// чем больше значение tuneCoefficient, тем меньше коэф разъебывается
	tuneCoefficient := float64(1)

	for i, v := range indexArray {

		// коэф по близости к началу
		result *= ((float64(len(wordsArray))-float64(v))/float64(len(wordsArray)))/tuneCoefficient + ensureValueIsMoreThanOne

		if i == len(indexArray)-1 {
			break
		}

		// коэф по распределенности слова
		result *= float64(indexArray[i+1]-v)/float64(len(wordsArray))/tuneCoefficient + ensureValueIsMoreThanOne
	}

	// коэф по количеству слова
	result *= float64(len(indexArray))/float64(len(wordsArray))/tuneCoefficient + ensureValueIsMoreThanOne

	return
}
