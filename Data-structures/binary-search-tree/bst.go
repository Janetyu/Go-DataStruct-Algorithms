package binary_search_tree

import (
	"fmt"
	"bytes"
	"strconv"
	"Go-DataStruct-Algorithms/Data-structures/stack"
	"reflect"
)

/**
这里的二分搜索树不包含重复元素，定义为：
每个节点的值：
	大于其左子树的所有节点的值
	小于其右子树的所有节点的值
每一棵子树也是二分搜索树
 */

type Node struct {
	e int
	left,right *Node
}

func NewNode(e int) *Node {
	return &Node{e,nil,nil}
}

type BST struct {
	root *Node
	size int
}

func NewBST() *BST {
	return &BST{nil,0}
}

func (b *BST)GetSize() int {
	return b.size
}

func (b *BST)IsEmpty() bool {
	return b.size == 0
}

// 向二分搜索树中添加新的元素
func (b *BST)Add(e int)  {

	//if b.root == nil {
	//	b.root = NewNode(e)
	//	b.size++
	//} else {
	//	b.add(b.root,e)
	//}

	b.root = b.add(b.root,e)
}

// 向以node为根的二分搜索树中插入元素e，递归算法
// 返回插入新节点后二分搜索树的根
func (b *BST)add(node *Node,e int) *Node {

	//if node.e == e {
	//	return
	//} else if node.left == nil && node.e > e{
	//	node.left = NewNode(e)
	//	b.size++
	//	return
	//} else if node.right == nil && node.e < e {
	//	node.right = NewNode(e)
	//	b.size++
	//	return
	//}
	//
	//if node.e > e {
	//	b.add(node.left,e)
	//} else { // node.e < e
	//	b.add(node.right,e)
	//}

	// 优化插入元素的方法
	if node == nil {
		b.size++
		return NewNode(e)
	}

	if node.e > e {
		node.left = b.add(node.left,e)
	} else if node.e < e {
		node.right = b.add(node.right,e)
	}

	return node
}

// 查询二分搜索树是否包含元素e
func (b *BST)Contains(e int) bool {
	return b.contains(b.root,e)
}

// 以node为根的二分搜索树中是否包含元素e，递归
func (b *BST)contains(node *Node,e int) bool {
	if node == nil {
		return false
	}

	if node.e == e {
		return true
	} else if node.e > e {
		return b.contains(node.left,e)
	} else { // node.e < e
		return b.contains(node.right,e)
	}
}

// 二分搜索树的前序遍历
func (b *BST)PreOrder()  {
	b.preOrder(b.root)
}

// 前序遍历以node为根的二分搜索树，递归算法
func (b *BST)preOrder(node *Node) {
	if node == nil {
		return
	}

	fmt.Println(node.e)
	b.preOrder(node.left)
	b.preOrder(node.right)
}

// 二分搜索树的前序遍历，非递归
func (b *BST)PreOrderNR()  {
	newstack := stack.CreatedefaultStack(reflect.Int)

	newstack.Push(b.root)
	for !newstack.IsEmpty() {
		cur := newstack.Pop()
		fmt.Println(cur.(*Node).e)

		// 后进先出，所以先压入右子树，再压入左子树
		if cur.(*Node).right != nil {
			newstack.Push(cur.(*Node).right)
		}
		if cur.(*Node).left != nil {
			newstack.Push(cur.(*Node).left)
		}
	}
}

// 二分搜索树的中序遍历
func (b *BST)InOrder()  {
	b.inOrder(b.root)
}

// 中序遍历以node为根的二分搜索树，递归算法
func (b *BST)inOrder(node *Node)  {
	if node == nil {
		return
	}

	b.inOrder(node.left)
	fmt.Println(node.e)
	b.inOrder(node.right)
}

// 二分搜索树的后序遍历
func (b *BST)PostOrder()  {
	b.inOrder(b.root)
}

// 后序遍历以node为根的二分搜索树，递归算法
func (b *BST)postOrder(node *Node)  {
	if node == nil {
		return
	}

	b.inOrder(node.left)
	b.inOrder(node.right)
	fmt.Println(node.e)
}

func (b *BST)String() string {
	var buffer bytes.Buffer
	b.generateBSTString(b.root,0,&buffer)
	return buffer.String()
}

// 生成以node为根节点，深度为depth的秒速二叉树的字符串
func (b *BST)generateBSTString(node *Node,depth int,buffer *bytes.Buffer)  {
	if node == nil {
		buffer.WriteString(b.generateDepthString(depth) + "null\n")
		return
	}

	buffer.WriteString(b.generateDepthString(depth) + strconv.Itoa(node.e) + "\n")
	b.generateBSTString(node.left, depth + 1, buffer)
	b.generateBSTString(node.right,depth + 1, buffer)
}

func (b *BST)generateDepthString(depth int) string {
	var buffer bytes.Buffer
	for i := 0; i < depth; i++ {
		buffer.WriteString("--")
	}
	return buffer.String()
}