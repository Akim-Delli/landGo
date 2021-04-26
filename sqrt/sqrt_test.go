package sqrt

import (
	"fmt"
	"testing"
)

func almostEquals(v1, v2 float64) bool {
	return Abs(v1-v2) <= 0.01
}

func TestSqrt(t *testing.T) {
	val, err := Sqrt(2)

	if err != nil {
		t.Fatalf("error in calculation %s", err)
	}

	if !almostEquals(val, 1.414214) {
		t.Fatalf("bad value - %f", val)
	}
}

type TestCase struct {
	value float64
	expected float64
}

func TestMany(t *testing.T) {
	testCases := []TestCase{
		{0.0, 0.0},
		{2.0, 1.414214},
		{9.0, 3.0},
	}


		for _, tc := range testCases {
			t.Run(fmt.Sprintf("%f", tc.value), func(t * testing.T) {
				out, err := Sqrt(tc.value)
				if err != nil {
					t.Fatal("error")
				}
				if !almostEquals(out , tc.expected){
					t.Fatalf("%f != %f", out, tc.expected)
				}
			})
		}
}