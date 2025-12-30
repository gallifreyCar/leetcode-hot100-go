# LeetCode Hot 100 - Go 语言题解

[![Go Version](https://img.shields.io/badge/Go-1.25-blue.svg)](https://golang.org)
[![LeetCode](https://img.shields.io/badge/LeetCode-Hot%20100-orange.svg)](https://leetcode.cn/studyplan/top-100-liked/)

LeetCode Hot 100 的 Go 语言实现，包含多种解法对比和详细注释。

## 📚 项目特点

✅ **标准函数名** - 与 LeetCode 官方完全一致，直接复制提交  
✅ **多解法对比** - 同一题目的不同解法独立管理  
✅ **完整测试** - 每个解法都有单元测试和性能测试  
✅ **易于调试** - IDE 中可直接运行/调试单个解法  

## 🗂️ 项目结构

```
problems/
├── 001_two_sum/              # ✅ 两数之和（已完成）
│   ├── solution.go
│   └── solution_test.go
├── 002_add_two_numbers/      # 📝 两数相加
│   ├── solution.go
│   └── solution_test.go
├── 003_longest_substring_without_repeating/  # 📝 无重复字符的最长子串
├── ...                       # 92 道题目
├── 049_group_anagrams/       # ✅ 字母异位词分组（多解法）
│   ├── v1/                   # 解法1：排序+哈希表
│   │   ├── solution.go
│   │   └── solution_test.go
│   ├── v2/                   # 解法2：计数+哈希表
│   │   ├── solution.go
│   │   └── solution_test.go
│   ├── v3/                   # 解法3：slices包
│   │   ├── solution.go
│   │   └── solution_test.go
│   └── solution_test.go      # 对比测试
└── 128_longest_consecutive/  # ✅ 最长连续序列（已完成）
    ├── solution.go
    └── solution_test.go
```

**说明**：
- 单一解法的题目：直接在题目目录下的 `solution.go`
- 多解法的题目：使用 `v1/`, `v2/`, `v3/` 子目录存放不同解法

## 🚀 使用方法

### 测试单个题目
```bash
go test ./problems/001_two_sum -v
```

### 测试某题的特定解法
```bash
go test ./problems/049_group_anagrams/v2 -v
```

### 测试所有题目
```bash
go test ./problems/... -v
```

### 性能对比
```bash
# 对比某题的所有解法性能
go test ./problems/049_group_anagrams -bench=. -benchmem

# 单个解法的性能测试
go test ./problems/049_group_anagrams/v2 -bench=. -benchmem
```

### 在 IDE 中调试
1. 打开 `v2/solution.go`
2. 在代码中设置断点
3. 打开 `v2/solution_test.go`
4. 右键测试函数 → **Debug Test**

## 📝 提交到 LeetCode

函数名完全标准，直接复制粘贴即可：

**进度：3 / 92 题**

<details>
<summary><b>📊 哈希表（3题）</b></summary>

| # | 题目 | 难度 | 状态 |
|---|------|------|------|
| 001 | [两数之和](problems/001_two_sum) | 简单 | ✅ |
| 049 | [字母异位词分组](problems/049_group_anagrams) | 中等 | ✅ (3种解法) |
| 128 | [最长连续序列](problems/128_longest_consecutive) | 中等 | ✅ |

</details>

<details>
<summary><b>🔄 双指针（4题）</b></summary>

| # | 题目 | 难度 | 状态 |
|---|------|------|------|
| 283 | [移动零](problems/283_move_zeroes) | 简单 | 📝 |
| 011 | [盛最多水的容器](problems/011_container_with_most_water) | 中等 | 📝 |
| 015 | [三数之和](problems/015_three_sum) | 中等 | 📝 |
| 042 | [接雨水](problems/042_trapping_rain_water) | 困难 | 📝 |

</details>

<details>
<summary><b>🪟 滑动窗口（5题）</b></summary>

| # | 题目 | 难度 | 状态 |
|---|------|------|------|
| 003 | [无重复字符的最长子串](problems/003_longest_substring_without_repeating) | 中等 | 📝 |
| 438 | [找到字符串中所有字母异位词](problems/438_find_all_anagrams) | 中等 | 📝 |
| 560 | [和为K的子数组](problems/560_subarray_sum_equals_k) | 中等 | 📝 |
| 239 | [滑动窗口最大值](problems/239_sliding_window_maximum) | 困难 | 📝 |
| 076 | [最小覆盖子串](problems/076_minimum_window_substring) | 困难 | 📝 |

</details>

<details>
<summary><b>📦 子数组（5题）</b></summary>

| # | 题目 | 难度 | 状态 |
|---|------|------|------|
| 053 | [最大子数组和](problems/053_maximum_subarray) | 中等 | 📝 |
| 056 | [合并区间](problems/056_merge_intervals) | 中等 | 📝 |
| 189 | [轮转数组](problems/189_rotate_array) | 中等 | 📝 |
| 238 | [除自身以外数组的乘积](problems/238_product_except_self) | 中等 | 📝 |
| 041 | [缺失的第一个正数](problems/041_first_missing_positive) | 困难 | 📝 |

</details>

<details>
<summary><b>🔲 矩阵（4题）</b></summary>

| # | 题目 | 难度 | 状态 |
|---|------|------|------|
| 073 | [矩阵置零](problems/073_set_matrix_zeroes) | 中等 | 📝 |
| 054 | [螺旋矩阵](problems/054_spiral_matrix) | 中等 | 📝 |
| 048 | [旋转图像](problems/048_rotate_image) | 中等 | 📝 |
| 240 | [搜索二维矩阵II](problems/240_search_2d_matrix_ii) | 中等 | 📝 |

</details>

<details>
<summary><b>🔗 链表（14题）</b></summary>

| # | 题目 | 难度 | 状态 |
|---|------|------|------|
| 160 | [相交链表](problems/160_intersection_of_two_linked_lists) | 简单 | 📝 |
| 206 | [反转链表](problems/206_reverse_linked_list) | 简单 | 📝 |
| 234 | [回文链表](problems/234_palindrome_linked_list) | 简单 | 📝 |
| 141 | [环形链表](problems/141_linked_list_cycle) | 简单 | 📝 |
| 142 | [环形链表II](problems/142_linked_list_cycle_ii) | 中等 | 📝 |
| 021 | [合并两个有序链表](problems/021_merge_two_sorted_lists) | 简单 | 📝 |
| 002 | [两数相加](problems/002_add_two_numbers) | 中等 | 📝 |
| 019 | [删除链表的倒数第N个结点](problems/019_remove_nth_node_from_end) | 中等 | 📝 |
| 024 | [两两交换链表中的节点](problems/024_swap_nodes_in_pairs) | 中等 | 📝 |
| 025 | [K个一组翻转链表](problems/025_reverse_nodes_in_k_group) | 困难 | 📝 |
| 138 | [随机链表的复制](problems/138_copy_list_with_random_pointer) | 中等 | 📝 |
| 148 | [排序链表](problems/148_sort_list) | 中等 | 📝 |
| 023 | [合并K个升序链表](problems/023_merge_k_sorted_lists) | 困难 | 📝 |
| 146 | [LRU缓存](problems/146_lru_cache) | 中等 | 📝 |

</details>

<details>
<summary><b>🌲 二叉树（15题）</b></summary>

| # | 题目 | 难度 | 状态 |
|---|------|------|------|
| 094 | [二叉树的中序遍历](problems/094_binary_tree_inorder_traversal) | 简单 | 📝 |
| 104 | [二叉树的最大深度](problems/104_maximum_depth_of_binary_tree) | 简单 | 📝 |
| 226 | [翻转二叉树](problems/226_invert_binary_tree) | 简单 | 📝 |
| 101 | [对称二叉树](problems/101_symmetric_tree) | 简单 | 📝 |
| 543 | [二叉树的直径](problems/543_diameter_of_binary_tree) | 简单 | 📝 |
| 102 | [二叉树的层序遍历](problems/102_binary_tree_level_order_traversal) | 中等 | 📝 |
| 108 | [将有序数组转换为二叉搜索树](problems/108_convert_sorted_array_to_bst) | 简单 | 📝 |
| 098 | [验证二叉搜索树](problems/098_validate_binary_search_tree) | 中等 | 📝 |
| 230 | [二叉搜索树中第K小的元素](problems/230_kth_smallest_element_in_bst) | 中等 | 📝 |
| 199 | [二叉树的右视图](problems/199_binary_tree_right_side_view) | 中等 | 📝 |
| 114 | [二叉树展开为链表](problems/114_flatten_binary_tree_to_linked_list) | 中等 | 📝 |
| 105 | [从前序与中序遍历序列构造二叉树](problems/105_construct_binary_tree) | 中等 | 📝 |
| 437 | [路径总和III](problems/437_path_sum_iii) | 中等 | 📝 |
| 236 | [二叉树的最近公共祖先](problems/236_lowest_common_ancestor) | 中等 | 📝 |
| 124 | [二叉树中的最大路径和](problems/124_binary_tree_maximum_path_sum) | 困难 | 📝 |

</details>

<details>
<summary><b>🗺️ 图论（4题）</b></summary>

| # | 题目 | 难度 | 状态 |
|---|------|------|------|
| 200 | [岛屿数量](problems/200_number_of_islands) | 中等 | 📝 |
| 994 | [腐烂的橘子](problems/994_rotting_oranges) | 中等 | 📝 |
| 207 | [课程表](problems/207_course_schedule) | 中等 | 📝 |
| 208 | [实现Trie](problems/208_implement_trie) | 中等 | 📝 |

</details>

<details>
<summary><b>🔙 回溯（8题）</b></summary>

| # | 题目 | 难度 | 状态 |
|---|------|------|------|
| 046 | [全排列](problems/046_permutations) | 中等 | 📝 |
| 078 | [子集](problems/078_subsets) | 中等 | 📝 |
| 017 | [电话号码的字母组合](problems/017_letter_combinations) | 中等 | 📝 |
| 039 | [组合总和](problems/039_combination_sum) | 中等 | 📝 |
| 022 | [括号生成](problems/022_generate_parentheses) | 中等 | 📝 |
| 079 | [单词搜索](problems/079_word_search) | 中等 | 📝 |
| 131 | [分割回文串](problems/131_palindrome_partitioning) | 中等 | 📝 |
| 051 | [N皇后](problems/051_n_queens) | 困难 | 📝 |

</details>

<details>
<summary><b>🔍 二分查找（6题）</b></summary>

| # | 题目 | 难度 | 状态 |
|---|------|------|------|
| 035 | [搜索插入位置](problems/035_search_insert_position) | 简单 | 📝 |
| 074 | [搜索二维矩阵](problems/074_search_2d_matrix) | 中等 | 📝 |
| 034 | [在排序数组中查找元素的第一个和最后一个位置](problems/034_find_first_and_last_position) | 中等 | 📝 |
| 033 | [搜索旋转排序数组](problems/033_search_in_rotated_sorted_array) | 中等 | 📝 |
| 153 | [寻找旋转排序数组中的最小值](problems/153_find_minimum_in_rotated_sorted_array) | 中等 | 📝 |
| 004 | [寻找两个正序数组的中位数](problems/004_median_of_two_sorted_arrays) | 困难 | 📝 |

</details>

<details>
<summary><b>📚 栈（5题）</b></summary>

| # | 题目 | 难度 | 状态 |
|---|------|------|------|
| 020 | [有效的括号](problems/020_valid_parentheses) | 简单 | 📝 |
| 155 | [最小栈](problems/155_min_stack) | 中等 | 📝 |
| 394 | [字符串解码](problems/394_decode_string) | 中等 | 📝 |
| 739 | [每日温度](problems/739_daily_temperatures) | 中等 | 📝 |
| 084 | [柱状图中最大的矩形](problems/084_largest_rectangle_in_histogram) | 困难 | 📝 |

</details>

<details>
<summary><b>🏔️ 堆（3题）</b></summary>

| # | 题目 | 难度 | 状态 |
|---|------|------|------|
| 215 | [数组中的第K个最大元素](problems/215_kth_largest_element) | 中等 | 📝 |
| 347 | [前K个高频元素](problems/347_top_k_frequent_elements) | 中等 | 📝 |
| 295 | [数据流的中位数](problems/295_find_median_from_data_stream) | 困难 | 📝 |

</details>

<details>
<summary><b>🎯 贪心（4题）</b></summary>

| # | 题目 | 难度 | 状态 |
|---|------|------|------|
| 121 | [买卖股票的最佳时机](problems/121_best_time_to_buy_and_sell_stock) | 简单 | 📝 |
| 055 | [跳跃游戏](problems/055_jump_game) | 中等 | 📝 |
| 045 | [跳跃游戏II](problems/045_jump_game_ii) | 中等 | 📝 |
| 763 | [划分字母区间](problems/763_partition_labels) | 中等 | 📝 |

</details>

<details>
<summary><b>💻 动态规划（10题）</b></summary>

| # | 题目 | 难度 | 状态 |
|---|------|------|------|
| 070 | [爬楼梯](problems/070_climbing_stairs) | 简单 | 📝 |
| 118 | [杨辉三角](problems/118_pascals_triangle) | 简单 | 📝 |
| 198 | [打家劫舍](problems/198_house_robber) | 中等 | 📝 |
| 279 | [完全平方数](problems/279_perfect_squares) | 中等 | 📝 |
| 322 | [零钱兑换](problems/322_coin_change) | 中等 | 📝 |
| 139 | [单词拆分](problems/139_word_break) | 中等 | 📝 |
| 300 | [最长递增子序列](problems/300_longest_increasing_subsequence) | 中等 | 📝 |
| 152 | [乘积最大子数组](problems/152_maximum_product_subarray) | 中等 | 📝 |
| 416 | [分割等和子集](problems/416_partition_equal_subset_sum) | 中等 | 📝 |
| 032 | [最长有效括号](problems/032_longest_valid_parentheses) | 困难 | 📝 |

</details>

<details>
<summary><b>🎲 多维动态规划（5题）</b></summary>

| # | 题目 | 难度 | 状态 |
|---|------|------|------|
| 062 | [不同路径](problems/062_unique_paths) | 中等 | 📝 |
| 064 | [最小路径和](problems/064_minimum_path_sum) | 中等 | 📝 |
| 005 | [最长回文子串](problems/005_longest_palindromic_substring) | 中等 | 📝 |
| 1143 | [最长公共子序列](problems/1143_longest_common_subsequence) | 中等 | 📝 |
| 072 | [编辑距离](problems/072_edit_distance) | 中等 | 📝 |

</details>
| # | 题目 | 难度 | 状态 |
|---|------|------|------|
| 003 | [无重复字符的最长子串](problems/003_longest_substring_without_repeating) | 中等 | 📝 |
| 438 | [找到字符串中所有字母异位词](problems/438_find_all_anagrams) | 中等 | 📝 |
| 560 | [和为K的子数组](problems/560_subarray_sum_equals_k) | 中等 | 📝 |
| 239 | [滑动窗口最大值](problems/239_sliding_window_maximum) | 困难 | 📝 |
| 076 | [最小覆盖子串](problems/076_minimum_window_substring) | 困难 | 📝 |

... 更多题目请查看目录

**图例**: ✅ 已完成 | 📝 待完成

## 💡 命名规范

- **包名**：`v1`, `v2`, `v3` (多解法) 或直接使用题目名
- **函数名**：与 LeetCode 完全一致（如 `TwoSum`, `GroupAnagrams`）
- **文件名**：`solution.go`, `solution_test.go`

## 🔧 环境要求

- Go 1.21+（使用 `slices` 包）
- 推荐使用 GoLand 或 VS Code

## 📖 学习资源

- [LeetCode Hot 100](https://leetcode.cn/studyplan/top-100-liked/)
- [Go 语言圣经](https://gopl-zh.github.io/)
- [代码随想录](https://programmercarl.com/)

## 📄 License

MIT License
