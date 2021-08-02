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
