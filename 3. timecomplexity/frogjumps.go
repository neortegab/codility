package timecomplexity

import "math"

func Solution(X, Y, D int) int {
	if X == Y {
		return 0
	}
	result := (float64(Y) - float64(X)) / float64(D)
	roundedResult := int(math.Round(result))

	if Y > D*roundedResult+X {
		return roundedResult + 1
	}
	return roundedResult
}
