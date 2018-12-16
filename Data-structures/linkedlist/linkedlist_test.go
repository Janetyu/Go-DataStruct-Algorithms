package linkedlist

import (
	"testing"
)

func Test_LinkedList(t *testing.T)  {
	linkedlist := CreateLinkedList()
	for i:=0 ; i < 5; i++ {
		linkedlist.AddFirst(i)
		t.Log(linkedlist.String())
	}

	linkedlist.Add(4,987)
	t.Log(linkedlist.String())

	linkedlist.Remove(4)
	t.Log(linkedlist.String())

	linkedlist.RemoveFirst()
	t.Log(linkedlist.String())

	linkedlist.RemoveLast()
	t.Log(linkedlist.String())
}
