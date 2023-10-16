package sorting

import (
	"sort"
)

func Distinct(A []int) int {
	if len(A) == 0 {
		return 0
	}
	sort.IntSlice.Sort(A)
	numDiff := 1
	for i := 0; i < len(A)-1; i++ {
		if A[i] != A[i+1] {
			numDiff++
		}
	}
	return numDiff

}

func MaxTriplProd(A []int) int {
	sort.IntSlice.Sort(A)
	pos1,pos2,pos3 := A[len(A)-1], A[len(A)-2], A[len(A)-3]
	neg1, neg2 := A[0], A[1]

	maxNegProd := neg1 * neg2 * pos1
	maxPosProd := pos1 * pos2 * pos3

	if maxNegProd > maxPosProd {
		return maxNegProd
	} else {
		return maxPosProd
	}
}
