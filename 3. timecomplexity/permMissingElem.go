package timecomplexity

import "sort"

func FindMissingElem(A []int) int {
	sort.IntSlice.Sort(A)
	missingNum := 1

	for _, v := range A {
		if v == missingNum {
			missingNum++
			continue
		}
		return missingNum
	}
	return missingNum
}
