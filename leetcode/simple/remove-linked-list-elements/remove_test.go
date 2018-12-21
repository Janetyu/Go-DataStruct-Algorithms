package remove_linked_list_elements

import "testing"

func Test_RemoveElements(t *testing.T)  {
	arr := []int{1,2,6,3,4,5,6,8,8,7,2,1,1}
	head := new(ListNode)
	head = head.CreateListNode(arr)
	t.Log(head.String())

	t.Logf("test removeElements: \n")
	res := removeElements(head,6)
	t.Log(res)

	t.Logf("test removeElements2: \n")
	res = removeElements2(head,8)
	t.Log(res)

	t.Logf("test removeElements3: \n")
	res = removeElements3(head,2)
	t.Log(res)

	t.Logf("test removeElements4: \n")
	res = removeElements4(head,1)
	t.Log(res)

	arr2 := []int{1,2,6,3,4,5,6}
	head = head.CreateListNode(arr2)
	t.Log(head.String())
	t.Logf("test removeElements5: \n")
	res = removeElements5(head,6,0)
	t.Log(res)

	t.Logf("test addElements: \n")
	res = addElements(head,6)
	t.Log(res)
}
