package num

import (
	"slices"
	"testing"
)

func TestEven(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	got := Even(numbers)
	want := []int{2, 4, 6, 8, 10}
	assertSliceEquality(t, got, want)
}

func TestOdd(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	got := Odd(numbers)
	want := []int{1, 3, 5, 7, 9}
	assertSliceEquality(t, got, want)
}

func TestPrime(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	got := Prime(numbers)
	want := []int{2, 3, 5, 7}
	assertSliceEquality(t, got, want)
}

func TestPrimeAndOdd(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	got := PrimeAndOdd(numbers)
	want := []int{3, 5, 7}
	assertSliceEquality(t, got, want)
}

func TestEvenAndMultipleOfFive(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	got := EvenAndMultipleOfFive(numbers)
	want := []int{10, 20}
	assertSliceEquality(t, got, want)
}

func TestOddAndMultipleOfThreeAndGreaterThanTen(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	got := OddAndMultipleOfThreeAndGreaterThanTen(numbers)
	want := []int{15}
	assertSliceEquality(t, got, want)
}

func assertSliceEquality(t testing.TB, got, want []int) {
	t.Helper()
	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
