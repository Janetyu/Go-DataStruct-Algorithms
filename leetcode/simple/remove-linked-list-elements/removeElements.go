package remove_linked_list_elements

/**
 * LeetCode T203
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }

	删除链表中等于给定值 val 的所有节点。

	示例:

	输入: 1->2->6->3->4->5->6, val = 6
	输出: 1->2->3->4->5
 */

// 不运用虚拟头结点
func removeElements(head *ListNode, val int) *ListNode {

	for head != nil && head.Val == val {
		delNode := head
		head = delNode.Next
		delNode.Next = nil
	}

	if head == nil {
		return nil
	}

	prev := head
	for prev.Next != nil {
		if prev.Next.Val == val {
			delNode := prev.Next
			prev.Next = delNode.Next
			delNode.Next = nil
		} else {
			prev = prev.Next
		}
	}

	return head
}

// 运用虚拟头结点
func removeElements2(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{Next:head}

	prev := dummyHead
	for prev.Next != nil {
		if prev.Next.Val == val {
			delNode := prev.Next
			prev.Next = delNode.Next
			delNode.Next = nil
		} else {
			prev = prev.Next
		}
	}

	return dummyHead.Next
}
