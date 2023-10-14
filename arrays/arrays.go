package arrays

import (
	"sort"
)

func CyclicRotation(A []int, K int) []int {
	if K == 0 || K == len(A) {
		return A
	}
	a := make([]int, len(A))
	copy(a, A)
	for i := 0; i <= K; i++ {
		for j := range A {
			if j == 0 {
				A[j] = a[len(A)-1]
				continue
			}
			A[j] = a[j-1]
		}
		copy(a, A)
	}
	return A
}

func OddOccurrencesInArray(A []int) int {
	if len(A) == 1 {
		return A[0]
	}
	sort.IntSlice.Sort(A)
	var oddOccurrence int
	for i := range A {
		if i == 0 && A[i] != A[i+1] {
			return A[i]
		} else if i > 0 && i < len(A)-1 && (A[i] != A[i-1] && A[i] != A[i+1]) {
			return A[i]
		} else if i == len(A)-1 && A[i] != A[i-1] {
			return A[i]
		}
	}
	return oddOccurrence
}
