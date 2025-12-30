package trappingrainwater

import "testing"

func TestTrap(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "示例1",
			height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			want:   6,
		},
		{
			name:   "示例2",
			height: []int{4, 2, 0, 3, 2, 5},
			want:   9,
		},
		{
			name:   "空数组",
			height: []int{},
			want:   0,
		},
		{
			name:   "单个元素",
			height: []int{5},
			want:   0,
		},
		{
			name:   "递增序列",
			height: []int{1, 2, 3, 4, 5},
			want:   0,
		},
		{
			name:   "山谷形状",
			height: []int{3, 0, 0, 2, 0, 4},
			want:   10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Trap(tt.height); got != tt.want {
				t.Errorf("Trap() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name+"_V2", func(t *testing.T) {
			if got := TrapV2(tt.height); got != tt.want {
				t.Errorf("TrapV2() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name+"_V3", func(t *testing.T) {
			if got := TrapV3(tt.height); got != tt.want {
				t.Errorf("TrapV3() = %v, want %v", got, tt.want)
			}
		})
	}
}
