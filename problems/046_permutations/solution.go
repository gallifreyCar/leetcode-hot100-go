package permutations

// 46. 全排列
// 难度：中等
// 标签：回溯
// 链接：https://leetcode.cn/problems/permutations/

// Permute 回溯
// 时间复杂度: O(n!)
// 空间复杂度: O(n)
func Permute(nums []int) [][]int {
	res := [][]int{}
	used := make([]bool, len(nums))
	path := []int{}

	var backtrack func()
	backtrack = func() {
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			path = append(path, nums[i])
			used[i] = true
			backtrack()
			path = path[:len(path)-1]
			used[i] = false
		}
	}

	backtrack()
	return res
}
