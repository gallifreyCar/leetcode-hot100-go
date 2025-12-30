package firstmissingpositive

import "testing"

func TestFirstMissingPositive(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "示例1",
			nums: []int{1, 2, 0},
			want: 3,
		},
		{
			name: "示例2",
			nums: []int{3, 4, -1, 1},
			want: 2,
		},
		{
			name: "示例3",
			nums: []int{7, 8, 9, 11, 12},
			want: 1,
		},
		{
			name: "连续序列",
			nums: []int{1, 2, 3, 4, 5},
			want: 6,
		},
		{
			name: "包含重复",
			nums: []int{1, 1, 2, 2},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := make([]int, len(tt.nums))
			copy(nums, tt.nums)
			if got := FirstMissingPositive(nums); got != tt.want {
				t.Errorf("FirstMissingPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}
