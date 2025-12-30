package v3

import (
	"reflect"
	"sort"
	"testing"
)

// 辅助函数：对结果进行排序以便比较
func sortResult(result [][]string) {
	for _, group := range result {
		sort.Strings(group)
	}
	sort.Slice(result, func(i, j int) bool {
		if len(result[i]) != len(result[j]) {
			return len(result[i]) < len(result[j])
		}
		return result[i][0] < result[j][0]
	})
}

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want [][]string
	}{
		{
			name: "示例1",
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			name: "示例2",
			strs: []string{""},
			want: [][]string{{""}},
		},
		{
			name: "示例3",
			strs: []string{"a"},
			want: [][]string{{"a"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 复制输入以避免修改
			input := make([]string, len(tt.strs))
			copy(input, tt.strs)

			got := GroupAnagrams(input)
			want := make([][]string, len(tt.want))
			for i := range tt.want {
				want[i] = make([]string, len(tt.want[i]))
				copy(want[i], tt.want[i])
			}

			sortResult(got)
			sortResult(want)

			if !reflect.DeepEqual(got, want) {
				t.Errorf("GroupAnagrams() = %v, want %v", got, want)
			}
		})
	}
}

// 基准测试
func BenchmarkGroupAnagrams(b *testing.B) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	for i := 0; i < b.N; i++ {
		GroupAnagrams(strs)
	}
}
