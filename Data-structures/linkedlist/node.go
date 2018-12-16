package linkedlist

import "fmt"

type Node struct {
	e interface{}
	next *Node
}

func CreateNode(e interface{},next *Node) *Node {
	return &Node{e,next}
}

func CreateDataNode(e interface{}) *Node {
	return CreateNode(e,&Node{})
}

func CreateNilNode() *Node {
	return CreateNode(nil,&Node{})
}

func (n *Node)String() string {
	return fmt.Sprint(n.e)
}
