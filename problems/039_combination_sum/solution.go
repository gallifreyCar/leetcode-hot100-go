package combinationsum

import "sort"

// 39. 组合总和
// 难度：中等
// 标签：回溯、数组
// 链接：https://leetcode.cn/problems/combination_sum/

// CombinationSum 回溯
// 时间复杂度: O(n * 2^n)
// 空间复杂度: O(target)
func CombinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	path := []int{}
	sort.Ints(candidates)

	var backtrack func(int, int)
	backtrack = func(start, sum int) {
		if sum == target {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		if sum > target {
			return
		}

		for i := start; i < len(candidates); i++ {
			path = append(path, candidates[i])
			backtrack(i, sum+candidates[i]) // 可以重复使用
			path = path[:len(path)-1]
		}
	}

	backtrack(0, 0)
	return res
}
