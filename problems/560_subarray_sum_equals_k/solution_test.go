package subarraysumequals

import "testing"

func TestSubarraySum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{"Example 1", []int{1, 1, 1}, 2, 2},
		{"Example 2", []int{1, 2, 3}, 3, 2},
		{"Negative numbers", []int{1, -1, 0}, 0, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubarraySum(tt.nums, tt.k); got != tt.want {
				t.Errorf("SubarraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
