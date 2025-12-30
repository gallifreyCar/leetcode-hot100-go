package rotatearray

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{
			name: "示例1",
			nums: []int{1, 2, 3, 4, 5, 6, 7},
			k:    3,
			want: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name: "示例2",
			nums: []int{-1, -100, 3, 99},
			k:    2,
			want: []int{3, 99, -1, -100},
		},
		{
			name: "k等于数组长度",
			nums: []int{1, 2, 3},
			k:    3,
			want: []int{1, 2, 3},
		},
		{
			name: "k大于数组长度",
			nums: []int{1, 2},
			k:    3,
			want: []int{2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := make([]int, len(tt.nums))
			copy(nums, tt.nums)
			Rotate(nums, tt.k)
			if !reflect.DeepEqual(nums, tt.want) {
				t.Errorf("Rotate() = %v, want %v", nums, tt.want)
			}
		})
		t.Run(tt.name+"_V2", func(t *testing.T) {
			nums := make([]int, len(tt.nums))
			copy(nums, tt.nums)
			RotateV2(nums, tt.k)
			if !reflect.DeepEqual(nums, tt.want) {
				t.Errorf("RotateV2() = %v, want %v", nums, tt.want)
			}
		})
	}
}
