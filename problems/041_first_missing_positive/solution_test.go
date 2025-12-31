package firstmissingpositive

import "testing"

func TestFirstMissingPositive(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example_1",
			nums: []int{1, 2, 0},
			want: 3,
		},
		{
			name: "Example_2",
			nums: []int{3, 4, -1, 1},
			want: 2,
		},
		{
			name: "Example_3",
			nums: []int{7, 8, 9, 11, 12},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := make([]int, len(tt.nums))
			copy(nums, tt.nums)
			if got := firstMissingPositive(nums); got != tt.want {
				t.Errorf("firstMissingPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}
