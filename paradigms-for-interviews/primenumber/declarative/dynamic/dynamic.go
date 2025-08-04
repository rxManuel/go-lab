package dynamic

import (
	"github.com/rxManuel/go-utils/fp-utils"
	"math"
)

var primeSet = map[int]struct{}{2: {}}

func hasNoDivisors(x int) bool {
    until := int(math.Sqrt(float64(x)))
    notDivisible := func(y int) bool { return !(x%y == 0)}
    return fputils.GenerateInclusiveIntRange(2, until).All(notDivisible)
}

func isCached(x int) bool {
	_, exists := primeSet[x]
	return exists
}

func IsPrime(number int) bool {

    if number > 1 && ( isCached(number) || hasNoDivisors(number) ) {
        primeSet[number] = struct{}{}
        return true
    }

    return false
}
