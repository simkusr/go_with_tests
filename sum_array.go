package sumarray

func SumArray(data []int) int {
	sum := 0
	for _, number := range data {
		sum += number
	}
	return sum
}
