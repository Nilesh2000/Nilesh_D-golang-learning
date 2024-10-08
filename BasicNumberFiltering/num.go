package num

import (
	"math"
)

func isEven(n int) bool {
	return n%2 == 0
}

func isOdd(n int) bool {
	return n%2 != 0
}

func isMultipleOfThree(n int) bool {
	return n%3 == 0
}

func isMultipleOfFive(n int) bool {
	return n%5 == 0
}

func isPrime(n int) bool {
	if n == 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func Even(n []int) []int {
	output := []int{}
	for _, num := range n {
		if isEven(num) {
			output = append(output, num)
		}
	}
	return output
}

func Odd(n []int) []int {
	output := []int{}
	for _, num := range n {
		if isOdd(num) {
			output = append(output, num)
		}
	}
	return output
}

func Prime(n []int) []int {
	output := []int{}
	for _, num := range n {
		if isPrime(num) {
			output = append(output, num)
		}
	}
	return output
}

func PrimeAndOdd(n []int) []int {
	output := []int{}
	for _, num := range n {
		if isPrime(num) && isOdd(num) {
			output = append(output, num)
		}
	}
	return output
}

func EvenAndMultipleOfFive(n []int) []int {
	output := []int{}
	for _, num := range n {
		if isEven(num) && isMultipleOfFive(num) {
			output = append(output, num)
		}
	}
	return output
}

func OddAndMultipleOfThreeAndGreaterThanTen(n []int) []int {
	output := []int{}
	for _, num := range n {
		if isOdd(num) && isMultipleOfThree(num) && num > 10 {
			output = append(output, num)
		}
	}
	return output
}
