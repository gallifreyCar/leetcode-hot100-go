package permutation_in_string

// CheckInclusion 固定窗口 + 字符计数数组
// 时间复杂度: O(n)
// 空间复杂度: O(1)
func CheckInclusion(s1 string, s2 string) bool {
	need := [26]int{}
	window := [26]int{}

	for i := 0; i < len(s1); i++ {
		need[s1[i]-'a']++
	}

	for start, end := 0, 0; end < len(s2); end++ {
		//固定窗口的板子（通用三步走）
		// 1. 加入新元素
		window[s2[end]-'a']++
		// 2. 窗口长度超过need固定长度，则收缩左边界
		if end-start+1 > len(s1) {
			window[s2[start]-'a']--
			start++
		}
		// 3. 窗口长度恰好等于need固定长度且符合need条件，处理结果
		if end-start+1 == len(s1) && need == window {
			return true
		}
	}
	return false
}
