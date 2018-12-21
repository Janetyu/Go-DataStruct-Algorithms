package remove_linked_list_elements

import (
	"bytes"
	"fmt"
	"strconv"
)

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

	// 去除中间有该值的情况
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

// 递归删除链表元素
func removeElements3(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	// 宏观语义：更小的问题：除去更小的链表中的值为val的元素
	res := removeElements3(head.Next,val)
	if head.Val == val {
		// 如果head结点的值为val，则去掉head，保留剩下更小的链表
		return res
	} else {
		head.Next = res
		return head
	}
}

// 递归实现尾部插入元素
func addElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return &ListNode{val,nil}
	} else {
		head.Next = addElements(head.Next,val)
		return head
	}
}

// 简化递归删除链表元素
func removeElements4(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	head.Next = removeElements4(head.Next,val)
	if head.Val == val {
		return head.Next
	} else {
		return head
	}

	// go语言没有三元运算符
	//return (head.Val == val)?head.Next:head
}

// 利用打印输出理解递归
func removeElements5(head *ListNode, val, depth int) *ListNode {

	depthString := genernateDepthString(depth)

	fmt.Print(depthString)
	fmt.Println("Call: remove " + strconv.Itoa(val) + " in " + head.String())
	if head == nil {
		fmt.Print(depthString)
		fmt.Println("Return: " + head.String())
		return nil
	}

	res := removeElements5(head.Next,val,depth+1)
	fmt.Print(depthString)
	fmt.Println("After remove " + strconv.Itoa(val) + ": " + res.String())

	ret := &ListNode{}
	if head.Val == val {
		ret = res
	} else {
		head.Next = res
		ret = head
	}
	fmt.Print(depthString)
	fmt.Println("Return: " + ret.String())
	return ret
}

func genernateDepthString(depth int) string {
	var buffer bytes.Buffer
	for i := 0; i < depth ; i++ {
		buffer.WriteString("--")
	}
	return buffer.String()
}


