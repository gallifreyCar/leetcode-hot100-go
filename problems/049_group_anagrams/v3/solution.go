package v3

import "slices"

// GroupAnagrams 使用新版本slices.Sort包
// 时间复杂度: O(n * k log k)
// 空间复杂度: O(n * k)
// 优点：使用Go 1.21+的新API，代码更简洁
func GroupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, str := range strs {
		s := []byte(str)
		slices.Sort(s)
		m[string(s)] = append(m[string(s)], str)
	}

	res := make([][]string, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}

	return res
}
