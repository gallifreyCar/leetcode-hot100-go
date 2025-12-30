package subsets

// 78. 子集
// 难度：中等
// 标签：回溯
// 链接：https://leetcode.cn/problems/subsets/

// Subsets 回溯
// 时间复杂度: O(2^n)
// 空间复杂度: O(n)
func Subsets(nums []int) [][]int {
	res := [][]int{}
	path := []int{}

	var backtrack func(int)
	backtrack = func(start int) {
		// 每个状态都是一个子集
		temp := make([]int, len(path))
		copy(temp, path)
		res = append(res, temp)

		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}

	backtrack(0)
	return res
}
