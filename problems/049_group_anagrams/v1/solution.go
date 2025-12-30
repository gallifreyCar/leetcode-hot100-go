package v1

import "sort"

// GroupAnagrams 排序+哈希表
// 时间复杂度: O(n * k log k), n是字符串数量，k是字符串平均长度
// 空间复杂度: O(n * k)
func GroupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)

	for _, str := range strs {
		// 排序
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		m[string(s)] = append(m[string(s)], str)
	}

	res := make([][]string, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
