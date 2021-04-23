package main

import (
	"fmt"
	"os"

	"github.com/Akim-Delli/landGo/contentType"
	"github.com/Akim-Delli/landGo/mascot"
	"github.com/Akim-Delli/landGo/serverkill"
	"github.com/Akim-Delli/landGo/wordCount"
	"rsc.io/quote"
)

func main() {
	fmt.Println(mascot.BestMascot())
	fmt.Println(quote.Go())

	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
	// fizzbuzz.Fizzbuzz()
	text := `
	Needles and Pins
	Needles and Pins
	Sew me a sail
	To catch me the wind
	`

	wordCount.WordCount(text)

	contentType, err := contentType.ContentType("http://www.google.com")
	fmt.Println(contentType, err)
	if err := serverkill.KillServer("/tmp/fdoesntexist"); err != nil {
		fmt.Fprintf(os.Stderr, "error:")
	}
}
