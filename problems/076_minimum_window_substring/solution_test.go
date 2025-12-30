package minimumwindowsubstring

import "testing"

func TestMinWindow(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want string
	}{
		{
			name: "示例1",
			s:    "ADOBECODEBANC",
			t:    "ABC",
			want: "BANC",
		},
		{
			name: "示例2",
			s:    "a",
			t:    "a",
			want: "a",
		},
		{
			name: "示例3",
			s:    "a",
			t:    "aa",
			want: "",
		},
		{
			name: "未找到",
			s:    "abc",
			t:    "xyz",
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinWindow(tt.s, tt.t); got != tt.want {
				t.Errorf("MinWindow() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name+"_V2", func(t *testing.T) {
			if got := MinWindowV2(tt.s, tt.t); got != tt.want {
				t.Errorf("MinWindowV2() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name+"_V3", func(t *testing.T) {
			if got := MinWindowV3(tt.s, tt.t); got != tt.want {
				t.Errorf("MinWindowV3() = %v, want %v", got, tt.want)
			}
		})
	}
}
