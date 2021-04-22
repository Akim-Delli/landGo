package wordCount

import (
	"fmt"
	"strings"
)

func WordCount(text string) map[string]int {
	words := strings.Fields(text)
	wordsCount := map[string]int{}

	for _, word := range words {
		wordsCount[strings.ToLower(word)]++
	}

	fmt.Println(wordsCount)
	return wordsCount

}
