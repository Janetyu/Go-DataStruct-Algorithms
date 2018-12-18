package remove_linked_list_elements

import (
	"bytes"
	"strconv"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func (listNode *ListNode)CreateListNode(arr []int) *ListNode {
	if arr == nil || len(arr) == 0 {
		panic("arr can not be empty")
	}

	listNode.Val = arr[0]
	cur := listNode
	for i := 1; i < len(arr); i++ {
		cur.Next = &ListNode{arr[i],nil}
		cur = cur.Next
	}

	return listNode
}

func (listNode *ListNode)String() string {
	var buffer bytes.Buffer
	cur := listNode
	for cur != nil {
		str := strconv.Itoa(cur.Val)
		buffer.WriteString(str + " -> ")
		cur = cur.Next
	}
	buffer.WriteString("NULL")
	return buffer.String()
}
