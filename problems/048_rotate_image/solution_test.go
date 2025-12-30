package rotateimage

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   [][]int
	}{
		{
			name: "示例1",
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			want: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		{
			name: "示例2",
			matrix: [][]int{
				{5, 1, 9, 11},
				{2, 4, 8, 10},
				{13, 3, 6, 7},
				{15, 14, 12, 16},
			},
			want: [][]int{
				{15, 13, 2, 5},
				{14, 3, 4, 1},
				{12, 6, 8, 9},
				{16, 7, 10, 11},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matrix := make([][]int, len(tt.matrix))
			for i := range tt.matrix {
				matrix[i] = make([]int, len(tt.matrix[i]))
				copy(matrix[i], tt.matrix[i])
			}
			Rotate(matrix)
			if !reflect.DeepEqual(matrix, tt.want) {
				t.Errorf("Rotate() = %v, want %v", matrix, tt.want)
			}
		})
	}
}
