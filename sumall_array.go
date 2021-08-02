package sumAll

func SumArray(data []int) int {
	sum := 0
	for _, number := range data {
		sum += number
	}
	return sum
}

func SumAll(data ...[]int) []int {
	lenghtOfNumbers := len(data)
	sums := make([]int, lenghtOfNumbers)
	for i, numbers := range data {
		sums[i] = SumArray(numbers)
	}
	return sums
}
