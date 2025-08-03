package concurrent

import "testing"

func TestIsPrimeImperativeConcurrent(t *testing.T) {
	t.Run("Negative", func(t *testing.T) {
		if got := IsPrime(-1); got != false {
			t.Errorf("IsPrime(-1) = %v; want false", got)
		}
	})

	t.Run("Zero", func(t *testing.T) {
		if got := IsPrime(0); got != false {
			t.Errorf("IsPrime(0) = %v; want false", got)
		}
	})

	t.Run("One", func(t *testing.T) {
		if got := IsPrime(1); got != false {
			t.Errorf("IsPrime(1) = %v; want false", got)
		}
	})

	t.Run("Two", func(t *testing.T) {
		if got := IsPrime(2); got != true {
			t.Errorf("IsPrime(2) = %v; want true", got)
		}
	})

	t.Run("Prime", func(t *testing.T) {
		if got := IsPrime(13); got != true {
			t.Errorf("IsPrime(13) = %v; want true", got)
		}
	})

	t.Run("NotPrime", func(t *testing.T) {
		if got := IsPrime(10); got != false {
			t.Errorf("IsPrime(10) = %v; want false", got)
		}
	})
	t.Run("NotPrime_10", func(t *testing.T) {
		if got := IsPrime(10); got != false {
			t.Errorf("IsPrime(10) = %v; want false", got)
		}
	})

	t.Run("Prime_101", func(t *testing.T) {
		if got := IsPrime(101); got != true {
			t.Errorf("IsPrime(101) = %v; want true", got)
		}
	})

	t.Run("NotPrime_121", func(t *testing.T) {
		if got := IsPrime(121); got != false { // 121 = 11 * 11
			t.Errorf("IsPrime(121) = %v; want false", got)
		}
	})

	t.Run("Prime_9973", func(t *testing.T) {
		if got := IsPrime(9973); got != true {
			t.Errorf("IsPrime(9973) = %v; want true", got)
		}
	})

	t.Run("NotPrime_10000", func(t *testing.T) {
		if got := IsPrime(10000); got != false {
			t.Errorf("IsPrime(10000) = %v; want false", got)
		}
	})

	t.Run("Prime_104729", func(t *testing.T) {
		if got := IsPrime(104729); got != true { // 10000th prime
			t.Errorf("IsPrime(104729) = %v; want true", got)
		}
	})
}
