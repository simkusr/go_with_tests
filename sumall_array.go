package sumAll

func SumAll(data ...[]int) []int {
	var result []int
	for i := 0; i < len(data); i++ {
		var sum int = 0
		for _, numb := range data[i] {
			sum += numb
		}
		result = append(result, sum)
	}
	return result
}
