package dynamic
// memoization

import "math"

// package level cache for prime numbers
// since in go there is no "Set"
// lets create a map with 0 bytes value
var primeSet = map[int]struct{}{} // or make(map[int]{})

func IsPrime(number int) bool {
	if number < 2 {
		return false
	}

	if number == 2 {
		return true
	} else {
		if _, exists := primeSet[number]; exists {
			return true
		}
		until := int(math.Sqrt(float64(number)))
		for n := 2; n <= until; n++ {
			if number%n == 0 {
				return false
			}
		}
	}
	primeSet[number] = struct{}{}
	return true
}
