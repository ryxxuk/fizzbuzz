package fizzbuzz

import (
	"strconv"
	"testing"
)

// Basic cases
func TestPrimeNumberReturnsUnchanged(t *testing.T) {
	var primeNumbers = [5]int{2, 7, 11, 13, 17}
	for _, num := range primeNumbers {
		result := fizzbuzz(num)
		if result != strconv.Itoa(num) {
			t.Errorf("Number was tampered with, got: %s, expected: %d.", result, num)
		}
	}
}

func TestMultipleOfThreeReturnsFizz(t *testing.T) {
	var multiplesOfThree = [5]int{3, 6, 9, 12, 18}
	for _, num := range multiplesOfThree {
		result := fizzbuzz(num)
		if result != "fizz" {
			t.Errorf("Multiple of 3 does not return fizz, got: %s.", result)
		}
	}
}

func TestMultipleOfFiveReturnsBuzz(t *testing.T) {
	var multiplesOfFive = [5]int{5, 10, 20, 25, 35}
	for _, num := range multiplesOfFive {
		result := fizzbuzz(num)
		if result != "buzz" {
			t.Errorf("Multiple of 5 does not return buzz, got: %s.", result)
		}
	}
}

func TestMultipleOfFiveAndThreeReturnsFizzBuzz(t *testing.T) {
	var multiplesOfFiveAndThree = [5]int{15, 30, 45, 60, 75}
	for _, num := range multiplesOfFiveAndThree {
		result := fizzbuzz(num)
		if result != "fizzbuzz" {
			t.Errorf("Multiple of 5 does not return fizzbuzz, got: %s.", result)
		}
	}
}

// Edge cases
func TestZeroReturnsFizzBuzz(t *testing.T) {
		result := fizzbuzz(0)
		if result != "fizzbuzz" {
			t.Errorf("0 does not return fizzbuzz, got: %s.", result)
		}
	
}

func TestHighNumberMultipleOfThreeReturnsFizz(t *testing.T) {
	result := fizzbuzz(2147483649)
	if result != "fizz" {
		t.Errorf("High number multiple of 5 does not return fizz, got: %s.", result)
	}

}

func TestHighNumberMultipleOfFiveReturnsBuzz(t *testing.T) {
	result := fizzbuzz(2147483650)
	if result != "buzz" {
		t.Errorf("High number multiple of 5 does not return buzz, got: %s.", result)
	}

}

func TestHighNumberMultipleOfThreeAndFiveReturnsFizzBuzz(t *testing.T) {
	result := fizzbuzz(2147483655)
	if result != "fizzbuzz" {
		t.Errorf("High number multiple of 5 and 3 does not return fizzbuzz, got: %s.", result)
	}
}

func TestHighNumberPrimesReturnsUnchanged(t *testing.T) {
	result := fizzbuzz(2147483659)
	if result != "2147483659" {
		t.Errorf("High number primes doesn't return unchanged, got: %s. expected: 2147483659", result)
	}

}

func TestNegativeNumberMultipleOfThreeReturnsFizz(t *testing.T) {
	result := fizzbuzz(-3)
	if result != "fizz" {
		t.Errorf("Negative multiple of 3 does not return fizz, got: %s.", result)
	}

}

func TestNegativeNumberMultipleOfFiveReturnsBuzz(t *testing.T) {
	result := fizzbuzz(-5)
	if result != "buzz" {
		t.Errorf("Negative multiple of 5 does not return buzz, got: %s.", result)
	}

}

func TestNegativeNumberMultipleOfThreeReturnsFizzBuzz(t *testing.T) {
	result := fizzbuzz(-15)
	if result != "fizzbuzz" {
		t.Errorf("Negative multiple of 5 and 3 does not return fizzbuzz, got: %s.", result)
	}
}

func TestNegativeNumberPrimesReturnsUnchanged(t *testing.T) {
	result := fizzbuzz(-1)
	if result != "-1" {
		t.Errorf("Negative prime number is not unchanged, got: %s. expected -1", result)
	}
}

// Range of values
func TestRangeOfValuesReturnsCorrectly(t *testing.T) {
	var valueRange = [15]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	var s = ""
	for _, num := range valueRange {
		s += fizzbuzz(num)
	}
	if s != "12fizz4buzzfizz78fizzbuzz11fizz1314fizzbuzz" {
		t.Errorf("Range of values does not return expected output, got: %s., expected 12fizz4buzzfizz78fizzbuzz11fizz1314fizzbuzz", s)
	}
}
