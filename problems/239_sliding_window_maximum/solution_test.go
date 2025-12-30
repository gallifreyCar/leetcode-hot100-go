package slidingwindowmaximum

import (
	"reflect"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{
			name: "示例1",
			nums: []int{1, 3, -1, -3, 5, 3, 6, 7},
			k:    3,
			want: []int{3, 3, 5, 5, 6, 7},
		},
		{
			name: "示例2",
			nums: []int{1},
			k:    1,
			want: []int{1},
		},
		{
			name: "k等于数组长度",
			nums: []int{1, 3, 2, 5},
			k:    4,
			want: []int{5},
		},
		{
			name: "递增序列",
			nums: []int{1, 2, 3, 4, 5},
			k:    2,
			want: []int{2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxSlidingWindow(tt.nums, tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxSlidingWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
