package v2

// GroupAnagrams 计数+哈希表
// 时间复杂度: O(n * k), n是字符串数量，k是字符串平均长度
// 空间复杂度: O(n * k)
// 优点：避免了排序，时间复杂度更优
func GroupAnagrams(strs []string) [][]string {
	// 用固定长度数组作为key
	m := make(map[[26]int][]string)
	for _, str := range strs {
		// 计数
		cnt := [26]int{}
		for _, s := range str {
			cnt[s-'a']++
		}
		m[cnt] = append(m[cnt], str)
	}

	res := make([][]string, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
