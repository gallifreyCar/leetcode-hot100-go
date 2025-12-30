package maximum_average_subarray_i

import "math"

// FindMaxAverage 固定窗口
// 时间复杂度: O(n)
// 空间复杂度: O(1)
func FindMaxAverage(nums []int, k int) float64 {
	var windowSum int
	res := math.Inf(-1)

	for start, end := 0, 0; end < len(nums); end++ {
		//固定窗口的板子（通用三步走）
		// 1. 加入新元素
		windowSum += nums[end]

		// 2. 窗口长度超过need固定长度，则收缩左边界
		if end-start+1 > k {
			windowSum -= nums[start]
			start++
		}
		// 3. 窗口长度恰好等于need固定长度且符合need条件，处理结果
		if end-start+1 == k {
			res = max(res, float64(windowSum)/float64(k))
		}
	}
	return res
}
