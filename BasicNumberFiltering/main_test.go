package main

import "testing"

func TestEven(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numbers = even(numbers)
	for _, num := range numbers {
		if num%2 != 0 {
			t.Errorf("%d is not even", num)
		}
	}
}

func TestOdd(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numbers = odd(numbers)
	for _, num := range numbers {
		if num%2 == 0 {
			t.Errorf("%d is not odd", num)
		}
	}
}
