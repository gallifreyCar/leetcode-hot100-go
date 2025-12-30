package longest_consecutive

// LongestConsecutive 哈希表解法
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func LongestConsecutive(nums []int) int {
	// 去重
	set := make(map[int]bool)
	for _, v := range nums {
		set[v] = true
	}

	var maxLen int
	for num := range set {
		// 不存在上一个数字，是起点，开始计算
		if _, ok := set[num-1]; !ok {
			curLen := 1
			curNum := num

			// 存在下一个数 +1
			for set[curNum+1] {
				curLen++
				curNum++
			}

			// 更新最长长度
			if curLen > maxLen {
				maxLen = curLen
			}

			// 有一个队列长度超过nums/2长度 提前返回
			if curLen*2 > len(nums) {
				break
			}
		}
	}

	return maxLen
}
