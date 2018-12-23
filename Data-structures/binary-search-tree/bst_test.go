package binary_search_tree

import (
	"testing"
	"math/rand"
	"math"
	"time"
)

func Test_BST(t *testing.T)  {
	bst := NewBST()
	//nums := []int{5,3,6,8,4,2}
	//for _,n := range nums {
	//	bst.Add(n)
	//}
	//
	//bst.PreOrder()
	//
	//bst.InOrder()
	//
	//bst.PostOrder()
	//
	//bst.PreOrderNR()
	//
	//t.Log("\n" + bst.String())

	// test removeMin
	n := 1000
	rand.Seed(time.Now().Unix())
	for i := 0 ; i < n ; i++ {
		rnum := rand.Intn(math.MaxInt16)
		bst.Add(rnum)
	}

	nums := []int{}
	for !bst.IsEmpty() {
		nums = append(nums,bst.RemoveMin())
	}

	t.Log(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			panic("Error.")
		}
	}
	t.Log("removeMin test completed.")


	// test removeMax
	for i := 0 ; i < n ; i++ {
		rnum := rand.Intn(math.MaxInt16)
		bst.Add(rnum)
	}

	nums = []int{}
	for !bst.IsEmpty() {
		nums = append(nums,bst.RemoveMax())
	}

	t.Log(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			panic("Error.")
		}
	}
	t.Log("removeMax test completed.")
}
