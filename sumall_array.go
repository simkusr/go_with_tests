package sumAll

func SumAll(arr1 []int) []int {
	var sum int = 0
	var result []int
	for _, numb := range arr1 {
		sum += numb
	}
	result = append(result, sum)
	return result
}
