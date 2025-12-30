package validparentheses

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "示例1",
			s:    "()",
			want: true,
		},
		{
			name: "示例2",
			s:    "()[]{}",
			want: true,
		},
		{
			name: "示例3",
			s:    "(]",
			want: false,
		},
		{
			name: "示例4",
			s:    "([)]",
			want: false,
		},
		{
			name: "嵌套括号",
			s:    "{[]}",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValid(tt.s); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
