package classical

import (
	"testing"
)

func assertEquals(t *testing.T, got, want bool) {
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestIsPrimeImperativeClassical(t *testing.T) {
	t.Run("is 1 prime?",
		func(t *testing.T) {
			got := IsPrime(1)
			want := false
			assertEquals(t, got, want)
		})

	t.Run("is 2 prime?",
		func(t *testing.T) {
			got := IsPrime(2)
			want := true
			assertEquals(t, got, want)
		})

	t.Run("is 3 prime?",
		func(t *testing.T) {
			got := IsPrime(3)
			want := true
			assertEquals(t, got, want)
		})

	t.Run("is 4 prime?",
		func(t *testing.T) {
			got := IsPrime(4)
			want := false
			assertEquals(t, got, want)
		})

	t.Run("is 17 prime?",
		func(t *testing.T) {
			got := IsPrime(17)
			want := true
			assertEquals(t, got, want)
		})

	t.Run("is 20 prime?",
		func(t *testing.T) {
			got := IsPrime(20)
			want := false
			assertEquals(t, got, want)
		})

	t.Run("is -1 prime?",
		func(t *testing.T) {
			got := IsPrime(-1)
			want := false
			assertEquals(t, got, want)
		})

	t.Run("is 307 prime?",
		func(t *testing.T) {
			got := IsPrime(307)
			want := true
			assertEquals(t, got, want)
		})

	t.Run("is 701 prime?",
		func(t *testing.T) {
			got := IsPrime(701)
			want := true
			assertEquals(t, got, want)
		})

}
