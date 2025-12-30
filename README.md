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
├── 001_two_sum/              # 两数之和
│   ├── solution.go
│   └── solution_test.go
├── 049_group_anagrams/       # 字母异位词分组
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
└── 128_longest_consecutive/  # 最长连续序列
    ├── solution.go
    └── solution_test.go
```

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

```bash
# 查看某个解法
cat problems/049_group_anagrams/v2/solution.go
```

复制 `GroupAnagrams` 函数，直接提交到 LeetCode！✅

## 🎯 题目列表

已完成 3 / 95 题

### 哈希表
| # | 题目 | 难度 | 状态 |
|---|------|------|------|
| 001 | [两数之和](problems/001_two_sum) | 简单 | ✅ |
| 049 | [字母异位词分组](problems/049_group_anagrams) | 中等 | ✅ |
| 128 | [最长连续序列](problems/128_longest_consecutive) | 中等 | ✅ |

### 双指针
| # | 题目 | 难度 | 状态 |
|---|------|------|------|
| 283 | [移动零](problems/283_move_zeroes) | 简单 | 📝 |
| 011 | [盛最多水的容器](problems/011_container_with_most_water) | 中等 | 📝 |
| 015 | [三数之和](problems/015_three_sum) | 中等 | 📝 |
| 042 | [接雨水](problems/042_trapping_rain_water) | 困难 | 📝 |

### 滑动窗口
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
