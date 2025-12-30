package threesum

import (
	"reflect"
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "示例1",
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "示例2",
			nums: []int{0, 1, 1},
			want: [][]int{},
		},
		{
			name: "示例3",
			nums: []int{0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
		{
			name: "全负数",
			nums: []int{-2, -3, -1, -4},
			want: [][]int{},
		},
		{
			name: "包含多组解",
			nums: []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6},
			want: [][]int{{-4, -2, 6}, {-4, 0, 4}, {-4, 1, 3}, {-4, 2, 2}, {-2, -2, 4}, {-2, 0, 2}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ThreeSum(tt.nums)
			// 排序结果以便比较
			for i := range got {
				sort.Ints(got[i])
			}
			sort.Slice(got, func(i, j int) bool {
				for k := 0; k < 3; k++ {
					if got[i][k] != got[j][k] {
						return got[i][k] < got[j][k]
					}
				}
				return false
			})
			for i := range tt.want {
				sort.Ints(tt.want[i])
			}
			sort.Slice(tt.want, func(i, j int) bool {
				for k := 0; k < 3; k++ {
					if tt.want[i][k] != tt.want[j][k] {
						return tt.want[i][k] < tt.want[j][k]
					}
				}
				return false
			})
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ThreeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
