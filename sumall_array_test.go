package sumAll

import (
	"reflect"
	"testing"
)

func TestSumArray(t *testing.T) {
	t.Run("Simple test", func(t *testing.T) {
		data := []int{1, 2, 3, 4, 5}
		got := SumArray(data)
		expected := 15
		if got != expected {
			t.Errorf("Expected %d but got %d while sent %v", expected, got, data)
		}
	})
}

// We need a new function called SumAll which will take a varying number of slices,
// returning a new slice containing the totals for each slice passed in.
// SumAll([]int{1,2}, []int{0,9}) would return []int{3, 9}
// SumAll([]int{1,1,1}) would return []int{3}

func TestSumAll(t *testing.T) {
	t.Run("Testing SumAll function with a single array", func(t *testing.T) {
		given := []int{1, 1, 1}
		expected := []int{3}
		got := SumAll(given)
		if !reflect.DeepEqual(expected, got) {
			t.Errorf("Expected: %v, but go: %v", expected, got)
		}
	})

	t.Run("Testing SumAll function with multiple arrays", func(t *testing.T) {
		givenOne := []int{1, 2}
		givenTwo := []int{0, 9}
		expected := []int{3, 9}
		got := SumAll(givenOne, givenTwo)
		if !reflect.DeepEqual(expected, got) {
			t.Errorf("Expected: %v, but go: %v", expected, got)
		}
	})

	t.Run("Testing SumAll function with 4 arrays", func(t *testing.T) {
		given1 := []int{1, 2}
		given2 := []int{0, 9}
		given3 := []int{1, 3, 1}
		given4 := []int{0, 1}
		expected := []int{3, 9, 5, 1}
		got := SumAll(given1, given2, given3, given4)
		if !reflect.DeepEqual(expected, got) {
			t.Errorf("Expected: %v, but go: %v", expected, got)
		}
	})
}
