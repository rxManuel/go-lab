package classical

func IsPrime(number int) bool {
	if number < 2 {
		return false
	}

	if number == 2 {
		return true
	} else {
		// Although iterating until sqrt(number) is more efficient,
		// we iterate until number / 2 to illustrate the concept:
		// no number greater than n/2 (except n itself) can be a proper divisor.
		for previous := 2; previous <= (number / 2); previous++ {
			if number%previous == 0 {
				// is divisible then false
				return false
			}
		}
	}
	return true
}
