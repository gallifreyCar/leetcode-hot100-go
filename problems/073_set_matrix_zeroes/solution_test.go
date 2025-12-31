package setmatrixzeroes

import (
	"reflect"
	"testing"
)

func TestSetZeroes(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   [][]int
	}{
		{
			name:   "Example_1",
			matrix: [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
			want:   [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}},
		},
		{
			name:   "Example_2",
			matrix: [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}},
			want:   [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matrix := make([][]int, len(tt.matrix))
			for i := range tt.matrix {
				matrix[i] = make([]int, len(tt.matrix[i]))
				copy(matrix[i], tt.matrix[i])
			}
			setZeroes(matrix)
			if !reflect.DeepEqual(matrix, tt.want) {
				t.Errorf("setZeroes() = %v, want %v", matrix, tt.want)
			}
		})
	}
}
