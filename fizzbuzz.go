package fizzbuzz

import "strconv"

func fizzbuzz(num int) string {
	if num % 5 == 0 && num % 3 == 0 {
		return "fizzbuzz"
	}
	if num % 3 == 0 {
		return "fizz"
	}
	if num % 5 == 0 {
		return "buzz"
	}

	return strconv.Itoa(num)
}