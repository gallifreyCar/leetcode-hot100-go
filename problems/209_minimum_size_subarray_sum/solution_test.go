package minimum_size_subarray_sum

import "testing"

func TestMinSubArrayLen(t *testing.T) {
	tests := []struct {
		name   string
		target int
		nums   []int
		expect int
	}{
		{"example1", 7, []int{2, 3, 1, 2, 4, 3}, 2},
		{"example2", 4, []int{1, 4, 4}, 1},
		{"example3", 11, []int{1, 1, 1, 1, 1, 1, 1, 1}, 0},
		{"single", 5, []int{5}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinSubArrayLen(tt.target, tt.nums); got != tt.expect {
				t.Errorf("MinSubArrayLen() = %v, want %v", got, tt.expect)
			}
		})
	}
}
