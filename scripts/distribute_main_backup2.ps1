
$files = @{
    "problems/142_linked_list_cycle_ii/solution.go" = @"
package linkedlistcycleii

type ListNode struct {
	Val  int
	Next *ListNode
}

// DetectCycle 快慢指针
// 结论：起点到环入口距离=相遇点到环入口距离
// 时间复杂度: O(n)
// 空间复杂度: O(1)
func DetectCycle(head *ListNode) *ListNode {
	fast := head
	slow := head
	var meetPoint *ListNode
	//找到相遇点
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			meetPoint = fast
			break
		}
	}
	if meetPoint == nil {
		return nil
	}

	// 再走一次
	cur := head
	for slow != cur {
		slow = slow.Next
		cur = cur.Next
	}
	return cur
}
"@

    "problems/002_add_two_numbers/solution.go" = @"
package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

// AddTwoNumbers 模拟加法
// 时间复杂度: O(max(m,n))
// 空间复杂度: O(1)
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	cnt := 0 //进位
	res := &ListNode{}
	cur := res
	cur1 := l1
	cur2 := l2

	for cur1 != nil || cur2 != nil || cnt != 0 {
		sum := 0
		if cur1 != nil {
			sum += cur1.Val
			cur1 = cur1.Next
		}
		if cur2 != nil {
			sum += cur2.Val
			cur2 = cur2.Next
		}
		sum += cnt
		cnt = sum / 10
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
	}

	return res.Next
}
"@

    "problems/019_remove_nth_node_from_end/solution.go" = @"
package removenthnodefromend

type ListNode struct {
	Val  int
	Next *ListNode
}

// RemoveNthFromEnd 快慢指针
// 时间复杂度: O(n)
// 空间复杂度: O(1)
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	fast := dummy
	slow := dummy
	//fast先走n步
	for n > 0 {
		fast = fast.Next
		n--
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next

	return dummy.Next
}
"@

    "problems/024_swap_nodes_in_pairs/solution.go" = @"
package swapnodesinpairs

type ListNode struct {
	Val  int
	Next *ListNode
}

// SwapPairs 迭代法
// 时间复杂度: O(n)
// 空间复杂度: O(1)
func SwapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		tmp := cur.Next           //1
		tmp2 := tmp.Next          //2
		tmp3 := tmp2.Next         //3
		cur.Next = tmp2           //0->2
		cur.Next.Next = tmp       //2->1
		cur.Next.Next.Next = tmp3 //1->3

		cur = cur.Next.Next //移动到下一对
	}

	return dummy.Next
}
"@

    "problems/025_reverse_nodes_in_k_group/solution.go" = @"
package reversenodesingroup

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

// ReverseKGroup K个一组翻转链表
// 时间复杂度: O(n)
// 空间复杂度: O(1)
func ReverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	start := head //翻转后头尾会反转
	cur := dummy  //用于链接反转后的链表
	end := dummy  //end从前一个节点出发
	for {
		// end先移动到k
		for i := 0; i < k; i++ {
			if end.Next == nil {
				return dummy.Next
			}
			end = end.Next
		}
		nextGroup := end.Next         //下一组
		end.Next = nil                //断开
		newHead := reverseList(start) //翻转链表
		cur.Next = newHead
		cur = start            //cur回到队尾帮忙链接
		end = start            //end回到队尾重新遍历
		start.Next = nextGroup //再链接下一组
		start = nextGroup
	}
}
"@

    "problems/138_copy_list_with_random_pointer/solution.go" = @"
package copylistwithrandompointer

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// CopyRandomList 哈希表
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func CopyRandomList(head *Node) *Node {
	//先生成对应的节点放入哈希表
	cur := head
	m := make(map[*Node]*Node)
	for cur != nil {
		m[cur] = &Node{Val: cur.Val}
		cur = cur.Next
	}

	//再赋值next和random
	cur = head
	for cur != nil {
		m[cur].Random = m[cur.Random]
		m[cur].Next = m[cur.Next]
		cur = cur.Next
	}

	return m[head]
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
