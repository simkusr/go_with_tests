package sumAll

import (
	"reflect"
	"testing"
)

// We need a new function called SumAll which will take a varying number of slices,
// returning a new slice containing the totals for each slice passed in.
// SumAll([]int{1,2}, []int{0,9}) would return []int{3, 9}
// SumAll([]int{1,1,1}) would return []int{3}

func TestSumAll(t *testing.T) {
	t.Run("Testing SumAll function with a single array", func(t *testing.T) {
		given := []int{1, 1, 1}
		expected := []int{3}
		got := SumAll(given)
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Expected: %v, but go : %v", expected, got)
		}
	})
}
