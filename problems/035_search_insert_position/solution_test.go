package searchinsertposition

import "testing"

func TestSearchInsertPosition(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "示例1: 找到目标值",
			nums:   []int{1, 3, 5, 6},
			target: 5,
			want:   2,
		},
		{
			name:   "示例2: 目标值不存在，应插入中间位置",
			nums:   []int{1, 3, 5, 6},
			target: 2,
			want:   1,
		},
		{
			name:   "示例3: 目标值不存在，应插入末尾",
			nums:   []int{1, 3, 5, 6},
			target: 7,
			want:   4,
		},
		{
			name:   "示例4: 单个元素，找到目标值",
			nums:   []int{1},
			target: 1,
			want:   0,
		},
		{
			name:   "示例5: 单个元素，插入到前面",
			nums:   []int{1},
			target: 0,
			want:   0,
		},
		{
			name:   "示例6: 单个元素，插入到后面",
			nums:   []int{1},
			target: 2,
			want:   1,
		},
		{
			name:   "示例7: 目标值在最前面",
			nums:   []int{1, 3, 5, 6},
			target: 0,
			want:   0,
		},
		{
			name:   "示例8: 找到第一个元素",
			nums:   []int{1, 3, 5, 6},
			target: 1,
			want:   0,
		},
		{
			name:   "示例9: 找到最后一个元素",
			nums:   []int{1, 3, 5, 6},
			target: 6,
			want:   3,
		},
		{
			name:   "示例10: 较大数组",
			nums:   []int{1, 3, 5, 7, 9, 11, 13, 15},
			target: 10,
			want:   5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchInsert(tt.nums, tt.target)
			if got != tt.want {
				t.Errorf("searchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}
