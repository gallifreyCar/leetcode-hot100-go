package palindromepartitioning

import (
	"reflect"
	"sort"
	"testing"
)

func TestPalindromePartitioning(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want [][]string
	}{
		{
			name: "示例1: aab",
			s:    "aab",
			want: [][]string{
				{"a", "a", "b"},
				{"aa", "b"},
			},
		},
		{
			name: "示例2: a",
			s:    "a",
			want: [][]string{
				{"a"},
			},
		},
		{
			name: "示例3: aba",
			s:    "aba",
			want: [][]string{
				{"a", "b", "a"},
				{"aba"},
			},
		},
		{
			name: "示例4: abac",
			s:    "abac",
			want: [][]string{
				{"a", "b", "a", "c"},
				{"a", "bab", "c"},
				{"aba", "c"},
			},
		},
		{
			name: "示例5: racecar",
			s:    "racecar",
			want: [][]string{
				{"r", "a", "c", "e", "c", "a", "r"},
				{"r", "a", "cec", "a", "r"},
				{"r", "aceca", "r"},
				{"racecar"},
			},
		},
		{
			name: "示例6: bb",
			s:    "bb",
			want: [][]string{
				{"b", "b"},
				{"bb"},
			},
		},
		{
			name: "空字符串",
			s:    "",
			want: [][]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := partition(tt.s)
			// 由于回溯可能产生不同顺序的结果，需要对结果进行排序后比较
			sortResult := func(res [][]string) {
				for i := range res {
					sort.Strings(res[i])
				}
				sort.Slice(res, func(i, j int) bool {
					if len(res[i]) != len(res[j]) {
						return len(res[i]) < len(res[j])
					}
					for k := range res[i] {
						if res[i][k] != res[j][k] {
							return res[i][k] < res[j][k]
						}
					}
					return false
				})
			}
			sortResult(got)
			sortResult(tt.want)
			
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}
