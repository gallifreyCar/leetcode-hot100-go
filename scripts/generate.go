package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Problem struct {
	ID         int
	Title      string
	TitleEN    string
	Difficulty string
	Tags       []string
}

var problems = []Problem{
	// 数组/哈希表
	{283, "移动零", "move_zeroes", "简单", []string{"双指针", "数组"}},
	{11, "盛最多水的容器", "container_with_most_water", "中等", []string{"双指针", "贪心"}},
	{15, "三数之和", "three_sum", "中等", []string{"双指针", "数组", "排序"}},
	{42, "接雨水", "trapping_rain_water", "困难", []string{"双指针", "动态规划", "栈"}},
	
	// 滑动窗口
	{3, "无重复字符的最长子串", "longest_substring_without_repeating", "中等", []string{"滑动窗口", "哈希表"}},
	{438, "找到字符串中所有字母异位词", "find_all_anagrams", "中等", []string{"滑动窗口", "哈希表"}},
	{560, "和为K的子数组", "subarray_sum_equals_k", "中等", []string{"前缀和", "哈希表"}},
	{239, "滑动窗口最大值", "sliding_window_maximum", "困难", []string{"滑动窗口", "单调队列"}},
	{76, "最小覆盖子串", "minimum_window_substring", "困难", []string{"滑动窗口", "哈希表"}},
	
	// 子数组
	{53, "最大子数组和", "maximum_subarray", "中等", []string{"动态规划", "分治"}},
	{56, "合并区间", "merge_intervals", "中等", []string{"数组", "排序"}},
	{189, "轮转数组", "rotate_array", "中等", []string{"数组"}},
	{238, "除自身以外数组的乘积", "product_except_self", "中等", []string{"数组", "前缀和"}},
	{41, "缺失的第一个正数", "first_missing_positive", "困难", []string{"数组", "哈希表"}},
	
	// 矩阵
	{73, "矩阵置零", "set_matrix_zeroes", "中等", []string{"数组", "矩阵"}},
	{54, "螺旋矩阵", "spiral_matrix", "中等", []string{"数组", "矩阵"}},
	{48, "旋转图像", "rotate_image", "中等", []string{"数组", "矩阵"}},
	{240, "搜索二维矩阵II", "search_2d_matrix_ii", "中等", []string{"数组", "二分查找", "矩阵"}},
	
	// 链表
	{160, "相交链表", "intersection_of_two_linked_lists", "简单", []string{"链表", "双指针"}},
	{206, "反转链表", "reverse_linked_list", "简单", []string{"链表", "递归"}},
	{234, "回文链表", "palindrome_linked_list", "简单", []string{"链表", "双指针"}},
	{141, "环形链表", "linked_list_cycle", "简单", []string{"链表", "双指针"}},
	{142, "环形链表II", "linked_list_cycle_ii", "中等", []string{"链表", "双指针"}},
	{21, "合并两个有序链表", "merge_two_sorted_lists", "简单", []string{"链表", "递归"}},
	{2, "两数相加", "add_two_numbers", "中等", []string{"链表", "递归"}},
	{19, "删除链表的倒数第N个结点", "remove_nth_node_from_end", "中等", []string{"链表", "双指针"}},
	{24, "两两交换链表中的节点", "swap_nodes_in_pairs", "中等", []string{"链表", "递归"}},
	{25, "K个一组翻转链表", "reverse_nodes_in_k_group", "困难", []string{"链表", "递归"}},
	{138, "随机链表的复制", "copy_list_with_random_pointer", "中等", []string{"链表", "哈希表"}},
	{148, "排序链表", "sort_list", "中等", []string{"链表", "归并排序"}},
	{23, "合并K个升序链表", "merge_k_sorted_lists", "困难", []string{"链表", "分治", "堆"}},
	{146, "LRU缓存", "lru_cache", "中等", []string{"设计", "哈希表", "链表"}},
	
	// 二叉树
	{94, "二叉树的中序遍历", "binary_tree_inorder_traversal", "简单", []string{"树", "栈"}},
	{104, "二叉树的最大深度", "maximum_depth_of_binary_tree", "简单", []string{"树", "深度优先搜索"}},
	{226, "翻转二叉树", "invert_binary_tree", "简单", []string{"树", "深度优先搜索"}},
	{101, "对称二叉树", "symmetric_tree", "简单", []string{"树", "深度优先搜索"}},
	{543, "二叉树的直径", "diameter_of_binary_tree", "简单", []string{"树", "深度优先搜索"}},
	{102, "二叉树的层序遍历", "binary_tree_level_order_traversal", "中等", []string{"树", "广度优先搜索"}},
	{108, "将有序数组转换为二叉搜索树", "convert_sorted_array_to_bst", "简单", []string{"树", "二叉搜索树"}},
	{98, "验证二叉搜索树", "validate_binary_search_tree", "中等", []string{"树", "二叉搜索树"}},
	{230, "二叉搜索树中第K小的元素", "kth_smallest_element_in_bst", "中等", []string{"树", "二叉搜索树"}},
	{199, "二叉树的右视图", "binary_tree_right_side_view", "中等", []string{"树", "广度优先搜索"}},
	{114, "二叉树展开为链表", "flatten_binary_tree_to_linked_list", "中等", []string{"树", "深度优先搜索"}},
	{105, "从前序与中序遍历序列构造二叉树", "construct_binary_tree", "中等", []string{"树", "递归"}},
	{437, "路径总和III", "path_sum_iii", "中等", []string{"树", "深度优先搜索"}},
	{236, "二叉树的最近公共祖先", "lowest_common_ancestor", "中等", []string{"树", "深度优先搜索"}},
	{124, "二叉树中的最大路径和", "binary_tree_maximum_path_sum", "困难", []string{"树", "深度优先搜索"}},
	
	// 图论
	{200, "岛屿数量", "number_of_islands", "中等", []string{"深度优先搜索", "广度优先搜索"}},
	{994, "腐烂的橘子", "rotting_oranges", "中等", []string{"广度优先搜索"}},
	{207, "课程表", "course_schedule", "中等", []string{"拓扑排序", "图"}},
	{208, "实现Trie", "implement_trie", "中等", []string{"设计", "字典树"}},
	
	// 回溯
	{46, "全排列", "permutations", "中等", []string{"回溯"}},
	{78, "子集", "subsets", "中等", []string{"回溯"}},
	{17, "电话号码的字母组合", "letter_combinations", "中等", []string{"回溯"}},
	{39, "组合总和", "combination_sum", "中等", []string{"回溯"}},
	{22, "括号生成", "generate_parentheses", "中等", []string{"回溯"}},
	{79, "单词搜索", "word_search", "中等", []string{"回溯", "矩阵"}},
	{131, "分割回文串", "palindrome_partitioning", "中等", []string{"回溯"}},
	{51, "N皇后", "n_queens", "困难", []string{"回溯"}},
	
	// 二分查找
	{35, "搜索插入位置", "search_insert_position", "简单", []string{"二分查找"}},
	{74, "搜索二维矩阵", "search_2d_matrix", "中等", []string{"二分查找"}},
	{34, "在排序数组中查找元素的第一个和最后一个位置", "find_first_and_last_position", "中等", []string{"二分查找"}},
	{33, "搜索旋转排序数组", "search_in_rotated_sorted_array", "中等", []string{"二分查找"}},
	{153, "寻找旋转排序数组中的最小值", "find_minimum_in_rotated_sorted_array", "中等", []string{"二分查找"}},
	{4, "寻找两个正序数组的中位数", "median_of_two_sorted_arrays", "困难", []string{"二分查找"}},
	
	// 栈
	{20, "有效的括号", "valid_parentheses", "简单", []string{"栈"}},
	{155, "最小栈", "min_stack", "中等", []string{"栈", "设计"}},
	{394, "字符串解码", "decode_string", "中等", []string{"栈"}},
	{739, "每日温度", "daily_temperatures", "中等", []string{"栈", "单调栈"}},
	{84, "柱状图中最大的矩形", "largest_rectangle_in_histogram", "困难", []string{"栈", "单调栈"}},
	
	// 堆
	{215, "数组中的第K个最大元素", "kth_largest_element", "中等", []string{"堆", "快速选择"}},
	{347, "前K个高频元素", "top_k_frequent_elements", "中等", []string{"堆", "哈希表"}},
	{295, "数据流的中位数", "find_median_from_data_stream", "困难", []string{"堆", "设计"}},
	
	// 贪心
	{121, "买卖股票的最佳时机", "best_time_to_buy_and_sell_stock", "简单", []string{"贪心", "动态规划"}},
	{55, "跳跃游戏", "jump_game", "中等", []string{"贪心", "动态规划"}},
	{45, "跳跃游戏II", "jump_game_ii", "中等", []string{"贪心"}},
	{763, "划分字母区间", "partition_labels", "中等", []string{"贪心"}},
	
	// 动态规划
	{70, "爬楼梯", "climbing_stairs", "简单", []string{"动态规划"}},
	{118, "杨辉三角", "pascals_triangle", "简单", []string{"动态规划"}},
	{198, "打家劫舍", "house_robber", "中等", []string{"动态规划"}},
	{279, "完全平方数", "perfect_squares", "中等", []string{"动态规划"}},
	{322, "零钱兑换", "coin_change", "中等", []string{"动态规划"}},
	{139, "单词拆分", "word_break", "中等", []string{"动态规划"}},
	{300, "最长递增子序列", "longest_increasing_subsequence", "中等", []string{"动态规划", "二分查找"}},
	{152, "乘积最大子数组", "maximum_product_subarray", "中等", []string{"动态规划"}},
	{416, "分割等和子集", "partition_equal_subset_sum", "中等", []string{"动态规划"}},
	{32, "最长有效括号", "longest_valid_parentheses", "困难", []string{"动态规划", "栈"}},
	
	// 多维动态规划
	{62, "不同路径", "unique_paths", "中等", []string{"动态规划"}},
	{64, "最小路径和", "minimum_path_sum", "中等", []string{"动态规划"}},
	{5, "最长回文子串", "longest_palindromic_substring", "中等", []string{"动态规划"}},
	{1143, "最长公共子序列", "longest_common_subsequence", "中等", []string{"动态规划"}},
	{72, "编辑距离", "edit_distance", "中等", []string{"动态规划"}},
}

func main() {
	for _, p := range problems {
		createProblem(p)
	}
	fmt.Println("✅ 已生成", len(problems), "道题目的模板")
}

func createProblem(p Problem) {
	dir := fmt.Sprintf("problems/%03d_%s", p.ID, p.TitleEN)
	os.MkdirAll(dir, 0755)
	
	// 生成函数名
	funcName := toCamelCase(p.TitleEN)
	
	// 创建 solution.go
	solutionPath := filepath.Join(dir, "solution.go")
	packageName := strings.ReplaceAll(p.TitleEN, "_", "")
	
	solutionContent := fmt.Sprintf(`package %s

// %d. %s
// 难度：%s
// 标签：%s
// 链接：https://leetcode.cn/problems/%s/
func %s() {
	// TODO: 实现代码
}
`, packageName, p.ID, p.Title, p.Difficulty, strings.Join(p.Tags, "、"), p.TitleEN, funcName)
	
	os.WriteFile(solutionPath, []byte(solutionContent), 0644)
	
	// 创建 solution_test.go
	testPath := filepath.Join(dir, "solution_test.go")
	testContent := fmt.Sprintf(`package %s

import "testing"

func Test%s(t *testing.T) {
	tests := []struct {
		name string
		// TODO: 添加测试用例字段
	}{
		{
			name: "示例1",
			// TODO: 填充测试数据
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: 调用函数并验证结果
		})
	}
}
`, packageName, funcName)
	
	os.WriteFile(testPath, []byte(testContent), 0644)
}

func toCamelCase(s string) string {
	parts := strings.Split(s, "_")
	for i := range parts {
		if len(parts[i]) > 0 {
			parts[i] = strings.ToUpper(parts[i][:1]) + parts[i][1:]
		}
	}
	return strings.Join(parts, "")
}
