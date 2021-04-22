package evenendedNumbers

import "fmt"


func EvenEnded(n int) bool {
	nAsStr := fmt.Sprintf("%d", n)
	return nAsStr[0] == nAsStr[len(nAsStr) - 1]
}