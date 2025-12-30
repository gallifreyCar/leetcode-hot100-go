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

| # | 题目 | 难度 | 解法数 | 标签 |
|---|------|------|--------|------|
| 001 | [两数之和](problems/001_two_sum) | 简单 | 1 | 哈希表 |
| 049 | [字母异位词分组](problems/049_group_anagrams) | 中等 | 3 | 哈希表、排序 |
| 128 | [最长连续序列](problems/128_longest_consecutive) | 中等 | 1 | 哈希表、并查集 |

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
