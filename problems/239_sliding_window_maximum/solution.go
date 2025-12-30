package slidingwindowmaximum

// 239. 滑动窗口最大值
// 难度：困难
// 标签：滑动窗口、单调队列
// 链接：https://leetcode.cn/problems/sliding_window_maximum/

// MaxSlidingWindow 单调队列
// 时间复杂度: O(n)
// 空间复杂度: O(k)
func MaxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return []int{}
	}

	res := make([]int, 0, len(nums)-k+1)
	deque := make([]int, 0) // 存储索引

	for i := 0; i < len(nums); i++ {
		// 移除窗口外的元素
		if len(deque) > 0 && deque[0] < i-k+1 {
			deque = deque[1:]
		}

		// 移除队尾比当前元素小的元素
		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[i] {
			deque = deque[:len(deque)-1]
		}

		deque = append(deque, i)

		// 当窗口形成时，记录最大值
		if i >= k-1 {
			res = append(res, nums[deque[0]])
		}
	}

	return res
}
