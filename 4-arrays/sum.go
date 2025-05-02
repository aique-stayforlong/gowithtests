package arrays

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(collections ...[]int) []int {
	var sums []int
	
	for _, numbers := range collections {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(collections ...[]int) []int {
	var sums []int

	for _, numbers := range collections {
		if len(numbers) > 0 {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		} else {
			sums = append(sums, 0)
		}
	}

	return sums
}