package permutations

import (
	"reflect"
	"sort"
	"testing"
)

func TestPermute(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "示例1",
			nums: []int{1, 2, 3},
			want: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
		{
			name: "示例2",
			nums: []int{0, 1},
			want: [][]int{{0, 1}, {1, 0}},
		},
		{
			name: "单个元素",
			nums: []int{1},
			want: [][]int{{1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Permute(tt.nums)
			// 排序以便比较
			for i := range got {
				sort.Ints(got[i])
			}
			sort.Slice(got, func(i, j int) bool {
				for k := 0; k < len(got[i]); k++ {
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
				for k := 0; k < len(tt.want[i]); k++ {
					if tt.want[i][k] != tt.want[j][k] {
						return tt.want[i][k] < tt.want[j][k]
					}
				}
				return false
			})
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Permute() = %v, want %v", got, tt.want)
			}
		})
	}
}
