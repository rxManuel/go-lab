package classical

import "math"

type intPredicate func(x int) bool

func rangeToFilter(start, end int) func(intPredicate) []int {
	return func(predicate intPredicate) []int {
		var filtered = []int{}
		for i := start; i <= end; i++ {
			if predicate(i) {
				filtered = append(filtered, i)
			}
		}
		return filtered
	}
}

func IsPrimeRaw(number int) bool {
	if number <= 1 {
		return false
	}
	if number == 2 {
		return true
	} else {
		isDivisible := func(x int) bool {
			return number%x == 0
		}

		until := int(math.Sqrt(float64(number)))
		divisors := rangeToFilter(2, until)(isDivisible)
		return len(divisors) == 0
	}
}
