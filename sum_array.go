package sumarray

func SumArray(data []int) int {
	sum := 0
	for i := 0; i < len(data); i++ {
		sum += data[i]
	}
	return sum
}
