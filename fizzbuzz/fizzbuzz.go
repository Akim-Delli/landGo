package fizzbuzz

import "fmt"

func Fizzbuzz() {
	for i := 1; i <= 20; i++ {
		switch {
		case isDivisibleBy(i, 15):
			fmt.Println("fizzbuzz")
		case isDivisibleBy(i, 3):
			fmt.Println("fizz")
		case isDivisibleBy(i, 5):
			fmt.Println("buzz")
		default:
			fmt.Println(i)

		}
	}
}

func isDivisibleBy(i, modulo int) bool {
	return i%modulo == 0
}
