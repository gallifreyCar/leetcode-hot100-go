package movezeroes

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"Example 1", []int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{"Example 2", []int{0}, []int{0}},
		{"All zeros", []int{0, 0, 0}, []int{0, 0, 0}},
		{"No zeros", []int{1, 2, 3}, []int{1, 2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MoveZeroes(tt.nums)
			if !reflect.DeepEqual(tt.nums, tt.want) {
				t.Errorf("MoveZeroes() = %v, want %v", tt.nums, tt.want)
			}
		})
	}
}
