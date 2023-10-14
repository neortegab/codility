package counting

func MinMissingPosNum(A []int) int {
	foundNums := make([]int, 1000000)
	for _, v := range A {
		if v > 0 {
			foundNums[v-1] = foundNums[v-1] + 1
		}
	}

	for i := range foundNums {
		if foundNums[i] == 0 {
			return i + 1
		}
	}
	return 1
}

func MaxCounters(N int, A []int) []int {
	result := make([]int, N)
	maxValue := 0
	augment := 0
	for _, v := range A {
		if v <= N {
			if augment > result[v-1] {
				result[v-1] = augment + 1
			} else {
				result[v-1] = result[v-1] + 1
			}
			if result[v-1] > maxValue {
				maxValue = result[v-1]
			}
		} else {
			augment = maxValue
		}
	}
	for i := range result {
		if result[i] < augment {
			result[i] = augment
		}
	}
	return result
}

func PermCheck(A []int) int {
	perm := make([]int, len(A))

	for _, v := range A {
		if v <= len(A) {
			permV := perm[v-1]
			perm[v-1] = permV + 1
		}
	}

	for _, v := range perm {
		if v == 0 || v > 1 {
			return 0
		}
	}

	return 1
}

func CountFrogJumps(X int, A []int) int {
	if len(A) == 1 {
		if A[0] == X {
			return 0
		} else {
			return -1
		}
	}

	var jumps []int

	for i := 0; i < X; i++ {
		jumps = append(jumps, -1)
	}

	for i, v := range A {
		if v <= X {
			if jumps[v-1] == -1 {
				jumps[v-1] = i
			}
		}
	}

	maxVal := -1
	for i := range jumps {
		if jumps[i] == -1 {
			return -1
		} else if jumps[i] > maxVal {
			maxVal = jumps[i]
		}
	}

	return maxVal
}
