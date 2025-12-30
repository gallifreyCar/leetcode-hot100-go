package max_consecutive_ones_iii

import "testing"

func TestLongestOnes(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		k      int
		expect int
	}{
		{"example1", []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2, 6},
		{"example2", []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3, 10},
		{"all_ones", []int{1, 1, 1, 1}, 0, 4},
		{"all_zeros", []int{0, 0, 0, 0}, 2, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestOnes(tt.nums, tt.k); got != tt.expect {
				t.Errorf("LongestOnes() = %v, want %v", got, tt.expect)
			}
		})
	}
}
