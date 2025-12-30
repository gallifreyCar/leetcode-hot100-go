package minimum_size_subarray_sum

// MinSubArrayLen 滑动窗口
// 时间复杂度: O(n)
// 空间复杂度: O(1)
func MinSubArrayLen(target int, nums []int) int {
	res := len(nums) + 1 //最后 res不可能大于len(nums)
	sum := 0
	for start, end := 0, 0; end < len(nums); end++ {
		//1.新增元素
		sum += nums[end]
		//2.收缩窗口
		for sum >= target {
			// 3.更新结果（这题要收缩前更新结果）
			res = min(res, end-start+1)
			sum -= nums[start]
			start++
		}
	}
	// 没找到答案
	if res == len(nums)+1 {
		return 0
	}
	return res
}
