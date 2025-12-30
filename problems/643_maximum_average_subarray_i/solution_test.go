package maximum_average_subarray_i

import (
	"math"
	"testing"
)

func TestFindMaxAverage(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		k      int
		expect float64
	}{
		{"example1", []int{1, 12, -5, -6, 50, 3}, 4, 12.75},
		{"example2", []int{5}, 1, 5.0},
		{"all_negative", []int{-1, -2, -3, -4}, 2, -1.5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindMaxAverage(tt.nums, tt.k)
			if math.Abs(got-tt.expect) > 1e-5 {
				t.Errorf("FindMaxAverage() = %v, want %v", got, tt.expect)
			}
		})
	}
}
