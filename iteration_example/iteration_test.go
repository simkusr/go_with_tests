package iterations

import (
	"fmt"
	"testing"
)

func TestIteration(t *testing.T) {
	got := Repeat("a", 6)
	expected := "aaaaaa"

	if got != expected {
		t.Errorf("expected %q but got %q", expected, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 6)
	}
}

func ExampleRepeat() {
	got := Repeat("a", 6)
	fmt.Println(got)
	// output: aaaaaa
}
