package longestsubstringwithoutrepeating

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "示例1",
			s:    "abcabcbb",
			want: 3,
		},
		{
			name: "示例2",
			s:    "bbbbb",
			want: 1,
		},
		{
			name: "示例3",
			s:    "pwwkew",
			want: 3,
		},
		{
			name: "空字符串",
			s:    "",
			want: 0,
		},
		{
			name: "单字符",
			s:    "a",
			want: 1,
		},
		{
			name: "全不同字符",
			s:    "abcdef",
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LengthOfLongestSubstring(tt.s); got != tt.want {
				t.Errorf("LengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name+"_V2", func(t *testing.T) {
			if got := LengthOfLongestSubstringV2(tt.s); got != tt.want {
				t.Errorf("LengthOfLongestSubstringV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
