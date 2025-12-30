package two_sum

// TwoSum 哈希表解法
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for loc, num := range nums {
		if v, ok := m[target-num]; ok {
			return []int{v, loc}
		}
		m[num] = loc
	}
	return nil
}
