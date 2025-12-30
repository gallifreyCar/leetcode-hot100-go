package spiralmatrix

import (
	"reflect"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   []int
	}{
		{
			name: "示例1",
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			want: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			name: "示例2",
			matrix: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
			},
			want: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			name: "单行",
			matrix: [][]int{
				{1, 2, 3, 4},
			},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "单列",
			matrix: [][]int{
				{1},
				{2},
				{3},
			},
			want: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SpiralOrder(tt.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SpiralOrder() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name+"_V2", func(t *testing.T) {
			if got := SpiralOrderV2(tt.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SpiralOrderV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
