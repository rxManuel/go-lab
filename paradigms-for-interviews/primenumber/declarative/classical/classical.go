package classical

import (
	"github.com/rxManuel/go-utils/fp-utils"
	"math"
)

func IsPrime(number int) bool {
	if number < 2  { return false }
	if number == 2 { return true }

	until := int(math.Sqrt(float64(number)))
	isDivisible := func(x int) bool { return number%x == 0 }
	return fputils.GenerateInclusiveIntRange(2, until).Filter(isDivisible).Count() == 0
}
