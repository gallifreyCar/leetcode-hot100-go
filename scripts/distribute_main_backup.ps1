
$files = @{
    "problems/240_search_2d_matrix_ii/solution.go" = @"
package search2dmatrixii

// 240. 搜索二维矩阵 II
// 难度：中等
// 标签：数组、二分查找、矩阵
// 链接：https://leetcode.cn/problems/search-a-2d-matrix-ii/
// SearchMatrix 从右上角开始搜索
// 时间复杂度: O(m+n)
// 空间复杂度: O(1)
func SearchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])

	j := n - 1
	i := 0
	for i < m && j >= 0 {
		//从右上角开始遍历
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] > target {
			j--
		} else if matrix[i][j] < target {
			i++
		}
	}
	return false
}
"@

    "problems/240_search_2d_matrix_ii/solution_test.go" = @"
package search2dmatrixii

import "testing"

func TestSearchMatrix(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		target int
		want   bool
	}{
		{
			"Example 1",
			[][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}},
			5,
			true,
		},
		{
			"Example 2",
			[][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}},
			20,
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchMatrix(tt.matrix, tt.target); got != tt.want {
				t.Errorf("SearchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
"@

    "problems/560_subarray_sum_equals_k/solution.go" = @"
package subarraysumequals

// 560. 和为 K 的子数组
// 难度：中等
// 标签：数组、哈希表、前缀和
// 链接：https://leetcode.cn/problems/subarray-sum-equals-k/
// SubarraySum 前缀和+哈希表优化（最优）
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func SubarraySum(nums []int, k int) int {
	// 前缀和+哈希表（两数之和的思想）
	// 优化掉前缀和数组，直接用一个变量记录当前前缀和
	m := make(map[int]int)
	m[0] = 1 // 前缀和为0出现一次
	res := 0
	prefixSum := 0

	for _, num := range nums {
		prefixSum += num
		if v, ok := m[prefixSum-k]; ok {
			res += v
		}
		m[prefixSum]++
	}

	return res
}
"@

    "problems/560_subarray_sum_equals_k/solution_test.go" = @"
package subarraysumequals

import "testing"

func TestSubarraySum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{\"Example 1\", []int{1, 1, 1}, 2, 2},
		{\"Example 2\", []int{1, 2, 3}, 3, 2},
		{\"Negative numbers\", []int{1, -1, 0}, 0, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubarraySum(tt.nums, tt.k); got != tt.want {
				t.Errorf(\"SubarraySum() = %v, want %v\", got, tt.want)
			}
		})
	}
}
"@

    "problems/160_intersection_of_two_linked_lists/solution.go" = @"
package intersectionoftwolinkedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

// GetIntersectionNode 双指针解法
// 时间复杂度: O(m+n)
// 空间复杂度: O(1)
func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	curA := headA
	curB := headB
	for curB != curA {
		if curA != nil {
			curA = curA.Next
		} else {
			curA = headB
		}

		if curB != nil {
			curB = curB.Next
		} else {
			curB = headA
		}
	}
	return curA
}
"@

    "problems/141_linked_list_cycle/solution.go" = @"
package linkedlistcycle

type ListNode struct {
	Val  int
	Next *ListNode
}

// HasCycle 快慢指针解法
// 时间复杂度: O(n)
// 空间复杂度: O(1)
func HasCycle(head *ListNode) bool {
	//快慢指针
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}
	return false
}
"@
}

foreach ($path in $files.Keys) {
    $content = $files[$path]
    $fullPath = $path
    $dir = Split-Path -Parent $fullPath
    if (!(Test-Path $dir)) {
        New-Item -ItemType Directory -Force -Path $dir | Out-Null
    }
    Set-Content -Path $fullPath -Value $content -Encoding UTF8
    Write-Host "Created/Updated $fullPath"
}
