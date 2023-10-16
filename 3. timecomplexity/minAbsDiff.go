package timecomplexity

import (
	"math"
)

func MinAbsDiff(A []int) int {
	sumT := 0
	for i := 0; i < len(A); i++ {
		sumT += A[i]
	}

	minDiff := 2000
	sL := 0
	for i := 0; i < len(A)-1; i++ {
		sL += A[i]
		sR := sumT - sL
		diff := int(math.Abs(float64(sL - sR)))
		if diff < minDiff {
			minDiff = diff
		}
	}

	return minDiff
}
