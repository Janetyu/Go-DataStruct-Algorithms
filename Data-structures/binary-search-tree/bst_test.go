package binary_search_tree

import "testing"

func Test_BST(t *testing.T)  {
	bst := NewBST()
	nums := []int{5,3,6,8,4,2}
	for _,n := range nums {
		bst.Add(n)
	}

	bst.PreOrder()

	bst.InOrder()

	bst.PostOrder()

	bst.PreOrderNR()

	t.Log("\n" + bst.String())
}
