package remove_linked_list_elements

import "testing"

func Test_RemoveElements(t *testing.T)  {
	arr := []int{1,2,6,3,4,5,6}
	head := new(ListNode)
	head = head.CreateListNode(arr)
	t.Log(head.String())

	t.Logf("test removeElements: \n")
	res := removeElements(head,6)
	t.Log(res)

	t.Logf("test removeElements2: \n")
	res = removeElements2(head,6)
	t.Log(res)
}
