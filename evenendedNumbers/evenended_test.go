package evenendedNumbers_test

import (
	"testing"

	"github.com/Akim-Delli/landGo/evenendedNumbers"
)

func TestEvenended(t *testing.T) {
	if !evenendedNumbers.EvenEnded(101) {
		t.Fatal("something when wrong")
	}
}

func TestEvenendedFalse(t *testing.T) {
	if evenendedNumbers.EvenEnded(1012) {
		t.Fatal("1012 is not even numbered when wrong")
	}
}
