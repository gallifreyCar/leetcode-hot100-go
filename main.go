package main

import (
	"fmt"

	two_sum "awesomeProject7/problems/001_two_sum"
	v1 "awesomeProject7/problems/049_group_anagrams/v1"
	v2 "awesomeProject7/problems/049_group_anagrams/v2"
	v3 "awesomeProject7/problems/049_group_anagrams/v3"
	longest_consecutive "awesomeProject7/problems/128_longest_consecutive"
)

func main() {
	// 快速测试用的main函数
	// 建议使用 go test 来运行单元测试

	fmt.Println("=== LeetCode 刷题项目 ===")

	// 示例1: Two Sum
	fmt.Println("\n测试 Two Sum:")
	result := two_sum.TwoSum([]int{2, 7, 11, 15}, 9)
	fmt.Printf("结果: %v\n", result)

	// 示例2: Group Anagrams (测试多个版本)
	fmt.Println("\n测试 Group Anagrams:")
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Printf("V1 结果: %v\n", v1.GroupAnagrams(strs))
	fmt.Printf("V2 结果: %v\n", v2.GroupAnagrams(strs))
	fmt.Printf("V3 结果: %v\n", v3.GroupAnagrams(strs))

	// 示例3: Longest Consecutive Sequence
	fmt.Println("\n测试 Longest Consecutive:")
	result2 := longest_consecutive.LongestConsecutive([]int{100, 4, 200, 1, 3, 2})
	fmt.Printf("结果: %v\n", result2)
}
