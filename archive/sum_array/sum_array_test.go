package sumarray

import "testing"

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
