package arrays

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
	countOfNums := make(map[int]int)
	
	for _, v := range A {
		val, isCounted := countOfNums[v]
		if !isCounted {
			countOfNums[v] = 1
		} else {
			countOfNums[v] = val + 1
		}
	}

	for num, count := range countOfNums {
		if count%2 != 0 {
			return num
		}
	}
	
	return 0
}
