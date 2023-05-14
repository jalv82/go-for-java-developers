package fizzbuzz

import (
	"strconv"
)

func FizzBuzz(input int) string {
	result := fizz(input) + buzz(input)
	if result != "" {
		return result
	}

	return strconv.Itoa(input)
}

func fizz(input int) string {
	if isDivisible(input, 3) {
		return "Fizz"
	}
	return ""
}

func buzz(input int) string {
	if isDivisible(input, 5) {
		return "Buzz"
	}
	return ""
}

func isDivisible(number, divisor int) bool {
	return number%divisor == 0
}
