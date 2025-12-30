package maximumsubarray

import "testing"

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "示例1",
			nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			want: 6,
		},
		{
			name: "示例2",
			nums: []int{1},
			want: 1,
		},
		{
			name: "示例3",
			nums: []int{5, 4, -1, 7, 8},
			want: 23,
		},
		{
			name: "全负数",
			nums: []int{-3, -2, -5, -1},
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxSubArray(tt.nums); got != tt.want {
				t.Errorf("MaxSubArray() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name+"_V2", func(t *testing.T) {
			if got := MaxSubArrayV2(tt.nums); got != tt.want {
				t.Errorf("MaxSubArrayV2() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name+"_V3", func(t *testing.T) {
			if got := MaxSubArrayV3(tt.nums); got != tt.want {
				t.Errorf("MaxSubArrayV3() = %v, want %v", got, tt.want)
			}
		})
	}
}
