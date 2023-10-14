package prefixsums

func PassingCars(A []int) int {
	numZeros := 0
	numPass := 0
	for i := range A {
		if A[i] == 0 {
			numZeros++
		} else {
			numPass += numZeros
		}
	}

	if numPass > 1000000000 {
		return -1
	}

	return numPass
}
