package concurrent

import "math"

func findDivisorInRange(number, start, end int, divisorFoundChannel chan bool) {
	for i := start; i <= end; i++ {
		if number%i == 0 {
			divisorFoundChannel <- true
			return
		}
	}
	divisorFoundChannel <- false
}

func IsPrime(number int) bool {
	if number <= 1 {
		return false
	}
	if number == 2 {
		return true
	}

	until := int(math.Sqrt(float64(number)))
	threadsQty := 2
	foundDivisorChan := make(chan bool, threadsQty)

	mid := until / threadsQty

	go findDivisorInRange(number, 2, mid, foundDivisorChan)
	go findDivisorInRange(number, mid+1, until, foundDivisorChan)

	for i := 0; i < threadsQty; i++ {
		hasDivisor := <-foundDivisorChan
		if hasDivisor {
			return false
		}
	}

	return true
}
