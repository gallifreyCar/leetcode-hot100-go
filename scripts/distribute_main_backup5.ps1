# 分发更多题解 - 链表和树

# 148_sort_list
$content = @'
package sortlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	//递归结束条件
	if head == nil || head.Next == nil {
		return head
	}
	//1.找中点
	mid := findMid(head) //偶数取左边，奇数取中间
	right := mid.Next
	mid.Next = nil //断开

	//2.归并
	left := sortList(head)
	right = sortList(right)

	//3.合并
	return merge(left, right)

}

func findMid(head *ListNode) *ListNode {
	slow, fast := head, head.Next //fast和slow起点不一样，确保slow停在左起点

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func merge(left *ListNode, right *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for left != nil && right != nil {
		if left.Val < right.Val {
			cur.Next = left
			left = left.Next
		} else {
			cur.Next = right
			right = right.Next
		}
		cur = cur.Next
	}
	if left != nil {
		cur.Next = left
	}
	if right != nil {
		cur.Next = right
	}
	return dummy.Next
}
'@
Set-Content -Path "problems/148_sort_list/solution.go" -Value $content -Encoding UTF8
Write-Host "Created/Updated problems/148_sort_list/solution.go"

$testContent = @'
package sortlist

import (
	"testing"
)

func TestSortList(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		want   []int
	}{
		{
			name:   "Example_1",
			values: []int{4, 2, 1, 3},
			want:   []int{1, 2, 3, 4},
		},
		{
			name:   "Example_2",
			values: []int{-1, 5, 3, 4, 0},
			want:   []int{-1, 0, 3, 4, 5},
		},
		{
			name:   "Empty",
			values: []int{},
			want:   []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := buildList(tt.values)
			got := sortList(head)
			gotValues := listToSlice(got)
			if !equal(gotValues, tt.want) {
				t.Errorf("sortList() = %v, want %v", gotValues, tt.want)
			}
		})
	}
}

func buildList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	head := &ListNode{Val: values[0]}
	cur := head
	for i := 1; i < len(values); i++ {
		cur.Next = &ListNode{Val: values[i]}
		cur = cur.Next
	}
	return head
}

func listToSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
'@
Set-Content -Path "problems/148_sort_list/solution_test.go" -Value $testContent -Encoding UTF8
Write-Host "Created/Updated problems/148_sort_list/solution_test.go"

# 023_merge_k_sorted_lists
$content = @'
package mergeksortedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	left := mergeKLists(lists[:len(lists)/2])
	right := mergeKLists(lists[len(lists)/2:])
	return merge(left, right)
}

func merge(left *ListNode, right *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for left != nil && right != nil {
		if left.Val < right.Val {
			cur.Next = left
			left = left.Next
		} else {
			cur.Next = right
			right = right.Next
		}
		cur = cur.Next
	}
	if left != nil {
		cur.Next = left
	}
	if right != nil {
		cur.Next = right
	}
	return dummy.Next
}
'@
Set-Content -Path "problems/023_merge_k_sorted_lists/solution.go" -Value $content -Encoding UTF8
Write-Host "Created/Updated problems/023_merge_k_sorted_lists/solution.go"

$testContent = @'
package mergeksortedlists

import "testing"

func TestMergeKLists(t *testing.T) {
	tests := []struct {
		name   string
		lists  [][]int
		want   []int
	}{
		{
			name:   "Example_1",
			lists:  [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}},
			want:   []int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		{
			name:   "Empty",
			lists:  [][]int{},
			want:   []int{},
		},
		{
			name:   "EmptyLists",
			lists:  [][]int{{}},
			want:   []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lists := make([]*ListNode, len(tt.lists))
			for i, vals := range tt.lists {
				lists[i] = buildList(vals)
			}
			got := mergeKLists(lists)
			gotValues := listToSlice(got)
			if !equal(gotValues, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", gotValues, tt.want)
			}
		})
	}
}

func buildList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	head := &ListNode{Val: values[0]}
	cur := head
	for i := 1; i < len(values); i++ {
		cur.Next = &ListNode{Val: values[i]}
		cur = cur.Next
	}
	return head
}

func listToSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
'@
Set-Content -Path "problems/023_merge_k_sorted_lists/solution_test.go" -Value $testContent -Encoding UTF8
Write-Host "Created/Updated problems/023_merge_k_sorted_lists/solution_test.go"

# 146_lru_cache
$content = @'
package lrucache

type Node struct {
	key, val   int
	prev, next *Node
}

type LRUCache struct {
	capacity int
	dummy    *Node //虚拟头节点
	tail     *Node //虚拟尾节点
	cache    map[int]*Node
}

func Constructor(capacity int) LRUCache {

	dummy := &Node{}
	tail := &Node{}
	dummy.next = tail
	tail.prev = dummy

	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		dummy:    dummy,
		tail:     tail,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.moveToHead(node)
		return node.val
	}
	return -1
}

// 新元素添加到头节点
func (this *LRUCache) addToHead(node *Node) {
	node.prev = this.dummy
	node.next = this.dummy.next
	this.dummy.next.prev = node
	this.dummy.next = node
}

// 移除节点
func (this *LRUCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// 移除尾节点(要返回给哈希表删除)
func (this *LRUCache) removeTail() *Node {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

// 移动到头节点
func (this *LRUCache) moveToHead(node *Node) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) Put(key int, value int) {
	//存在
	if node, ok := this.cache[key]; ok {
		node.val = value
		this.moveToHead(node)
		return
	}

	node := &Node{key: key, val: value}
	this.cache[key] = node
	this.addToHead(node)

	if len(this.cache) > this.capacity {
		removeTail := this.removeTail()
		delete(this.cache, removeTail.key)
	}

}
'@
Set-Content -Path "problems/146_lru_cache/solution.go" -Value $content -Encoding UTF8
Write-Host "Created/Updated problems/146_lru_cache/solution.go"

$testContent = @'
package lrucache

import "testing"

func TestLRUCache(t *testing.T) {
	t.Run("Example_1", func(t *testing.T) {
		cache := Constructor(2)
		cache.Put(1, 1)
		cache.Put(2, 2)
		if got := cache.Get(1); got != 1 {
			t.Errorf("Get(1) = %v, want 1", got)
		}
		cache.Put(3, 3) // evicts key 2
		if got := cache.Get(2); got != -1 {
			t.Errorf("Get(2) = %v, want -1", got)
		}
		cache.Put(4, 4) // evicts key 1
		if got := cache.Get(1); got != -1 {
			t.Errorf("Get(1) = %v, want -1", got)
		}
		if got := cache.Get(3); got != 3 {
			t.Errorf("Get(3) = %v, want 3", got)
		}
		if got := cache.Get(4); got != 4 {
			t.Errorf("Get(4) = %v, want 4", got)
		}
	})
}
'@
Set-Content -Path "problems/146_lru_cache/solution_test.go" -Value $testContent -Encoding UTF8
Write-Host "Created/Updated problems/146_lru_cache/solution_test.go"

Write-Host "`nAll files created/updated successfully!"
