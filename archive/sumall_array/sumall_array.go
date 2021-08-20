package sumAll

func SumArray(data []int) int {
	sum := 0
	for _, number := range data {
		sum += number
	}
	return sum
}

func SumAll(data ...[]int) []int {
	var sums []int
	for _, numbers := range data {
		sums = append(sums, SumArray(numbers))
	}
	return sums
}

func SumAllTails(data ...[]int) []int {
	var sum []int
	for _, numbers := range data {
		if len(numbers) == 0 {
			sum = append(sum, 0)
		} else {
			tail := numbers[1:]
			sum = append(sum, SumArray(tail))
		}
	}
	return sum
}
