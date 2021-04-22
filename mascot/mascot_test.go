package mascot_test

import (
	"testing"

	"github.com/Akim-Delli/landGo/mascot"
)

func TestMascot(t *testing.T) {
	if mascot.BestMascot() != "Tux" {
		t.Fatal("Go Gopher")
	}

}
