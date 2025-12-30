package range_sum_query_immutable

import "testing"

func TestNumArray(t *testing.T) {
	nums := []int{-2, 0, 3, -5, 2, -1}
	obj := Constructor(nums)

	tests := []struct {
		name   string
		left   int
		right  int
		expect int
	}{
		{"range1", 0, 2, 1},
		{"range2", 2, 5, -1},
		{"range3", 0, 5, -3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := obj.SumRange(tt.left, tt.right); got != tt.expect {
				t.Errorf("SumRange() = %v, want %v", got, tt.expect)
			}
		})
	}
}
