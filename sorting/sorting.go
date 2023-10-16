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
	triplet := make([]int, 3)

    for i := range A {
		if triplet[0] == 0 {
			triplet[0] = A[i]
		} else if triplet[1] == 0 {
			triplet[1] = A[i]
		} else if triplet[2] == 0{
			triplet[2] = A[i]
		} else if A[i] > triplet[0] {
			triplet[0] = A[i]
		} else if A[i] > triplet[1] {
			triplet[1] = A[i]
		} else if A[i] > triplet[2] {
			triplet[2] = A[i]
		}
    }

	return triplet[0] * triplet[1] * triplet[2]
}
