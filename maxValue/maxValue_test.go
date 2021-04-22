package maxValue_test

import (
	"testing"

	"github.com/Akim-Delli/landGo/maxValue"
)

func TestMaxValue(t *testing.T) {
	MAXIMUMVALUE := 48
	nums := []int{16, 8, MAXIMUMVALUE, 23, 15}

	if maxValue.MaxValue(nums) != MAXIMUMVALUE {
		t.Fatal("MaxValue doesn't return the max Value")
	}
}
