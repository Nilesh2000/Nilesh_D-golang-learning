package main

import (
	"fmt"
	"math"
)

type numbers []int

func isOdd(n int) bool {
	return n%2 != 0
}

func isEven(n int) bool {
	return n%2 == 0
}

func isMultipleOfThree(n int) bool {
	return n%3 == 0
}

func isMultipleOfFive(n int) bool {
	return n%5 == 0
}

func even(n numbers) numbers {
	output := numbers{}
	for _, num := range n {
		if isEven(num) {
			output = append(output, num)
		}
	}
	return output
}

func odd(n numbers) numbers {
	output := numbers{}
	for _, num := range n {
		if isOdd(num) {
			output = append(output, num)
		}
	}
	return output
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

func prime(n numbers) numbers {
	output := numbers{}
	for _, num := range n {
		if isPrime(num) {
			output = append(output, num)
		}
	}
	return output
}

func primeAndOdd(n numbers) numbers {
	output := numbers{}
	for _, num := range n {
		if isPrime(num) && isOdd(num) {
			output = append(output, num)
		}
	}
	return output
}

func evenAndMultipleOfFive(n numbers) numbers {
	output := numbers{}
	for _, num := range n {
		if isEven(num) && isMultipleOfFive(num) {
			output = append(output, num)
		}
	}
	return output
}

func oddAndMultipleOfThreeAndGreaterThanTen(n numbers) numbers {
	output := numbers{}
	for _, num := range n {
		if isOdd(num) && isMultipleOfThree(num) && num > 10 {
			output = append(output, num)
		}
	}
	return output
}

func (n numbers) print() {
	for _, num := range n {
		fmt.Print(num, " ")
	}
	fmt.Println()
}

func main() {
	nums := numbers{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	evenNums := even(nums)
	evenNums.print()

	oddNums := odd(nums)
	oddNums.print()

	primeNums := prime(nums)
	primeNums.print()

	primeAndOddNums := primeAndOdd(nums)
	primeAndOddNums.print()

	nums = numbers{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	evenAndMultipleOfFiveNums := evenAndMultipleOfFive(nums)
	evenAndMultipleOfFiveNums.print()

	oddAndMultipleOfThreeAndGreaterThanTenNums := oddAndMultipleOfThreeAndGreaterThanTen(nums)
	oddAndMultipleOfThreeAndGreaterThanTenNums.print()
}
