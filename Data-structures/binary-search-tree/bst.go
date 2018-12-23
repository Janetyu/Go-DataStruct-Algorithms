package binary_search_tree

import (
	"fmt"
	"bytes"
	"strconv"
	"reflect"

	"Go-DataStruct-Algorithms/Data-structures/stack"
	"Go-DataStruct-Algorithms/Data-structures/queue"
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
	// 利用了栈这个结构
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

// 层序遍历二分搜索树（广度优先遍历）
func (b *BST)LevelOrder()  {
	// 利用了队列这种数据结构
	queue := queue.CreatedefaultQueue(reflect.Int)

	// 根节点先入队
	queue.Enqueue(b.root)
	for !queue.IsEmpty() {
		cur := queue.Dequeue()
		fmt.Println(cur.(*Node).e)

		// 队列先进先出，所以左子树先入队，右子树后入队
		if cur.(*Node).left != nil {
			queue.Enqueue(cur.(*Node).left)
		}
		if cur.(*Node).right != nil {
			queue.Enqueue(cur.(*Node).right)
		}
	}
}

// 寻找二分搜索树的最小元素
func (b *BST)Minimum() int {
	if b.size == 0 {
		panic("BST is empty!")
	}

	return b.minimum(b.root).e
}

// 返回以node为根的二分搜索树的最小值所在的节点
func (b *BST)minimum(node *Node) *Node {
	if node.left == nil {
		return node
	}
	return b.minimum(node.left)
}

// 非递归实现寻找二分搜索树的最小元素
func (b *BST)MinimumNR() int {
	if b.size == 0 {
		panic("BST is empty!")
	}

	cur := b.root
	if cur.left != nil {
		cur = cur.left
	}

	// cur.left == nil
	return cur.e
}

// 寻找二分搜索树的最大元素
func (b *BST)Maximum() int {
	if b.size == 0 {
		panic("BST is empty!")
	}

	return b.maximum(b.root).e
}

// 返回以node为根的二分搜索树的最大值所在的节点
func (b *BST)maximum(node *Node) *Node {
	if node.right == nil {
		return node
	}
	return b.maximum(node.right)
}

// 非递归实现寻找二分搜索树的最大元素
func (b *BST)MaximumNR() int {
	if b.size == 0 {
		panic("BST is empty!")
	}

	cur := b.root
	if cur.right != nil {
		cur = cur.right
	}

	// cur.left == nil
	return cur.e
}

// 从二分搜索树中删除最小值所在的节点，并返回最小值
func (b *BST)RemoveMin() int {
	min := b.Minimum()
	b.root = b.removeMin(b.root)
	return min
}

// 删除掉以node为根的二分搜索树中的最小节点
// 返回删除节点中后的新的二分搜索树的根
func (b *BST)removeMin(node *Node) *Node {
	// 终止条件，左子树为空
	if node.left == nil {
		rightNode := node.right
		node.right = nil
		b.size--
		return rightNode
	}

	// 不管node.right是否为空，若最小值被删除，则最小值中的右子树变为上一节点的左子树
	node.left = b.removeMin(node.left)
	return node
}

// 从二分搜索树中删除最大值所在的节点，并返回最大值
func (b *BST)RemoveMax() int {
	max := b.Maximum()
	b.root = b.removeMax(b.root)
	return max
}

// 删除掉以node为根的二分搜索树中的最大节点
// 返回删除节点中后的新的二分搜索树的根
func (b *BST)removeMax(node *Node) *Node {
	// 终止条件，右子树为空
	if node.right == nil {
		leftNode := node.left
		node.left = nil
		b.size--
		return leftNode
	}

	// 不管node.left是否为空，若最大值被删除，则最大值中的左子树变为上一节点的右子树
	node.right = b.removeMax(node.right)
	return node
}

// 从二分搜索树中删除元素为e的节点
func (b *BST)Remove(e int)  {
	b.root = b.remove(b.root,e)
}

// 删除以node为根的二分搜索树中值为e的节点，递归算法
// 返回删除节点后的新的二分搜索树的根
func (b *BST)remove(node *Node, e int) *Node {
	// 如果节点为空，则什么都不用干
	if node == nil {
		return nil
	}

	if node.e > e { // 先判断小于的情况，去左子树进行搜索
		node.left = b.remove(node.left,e)
		return node
	} else if node.e < e { // 先判断大于的情况，去右子树进行搜索
		node.right = b.remove(node.right,e)
		return node
	} else { // 最后判断等于的情况，分只有左子树，只有右子树，有左右子树的情况
		// 删除的节点有右子树,左子树为空
		if node.left == nil {
			rightNode := node.right
			node.right = nil
			b.size--
			return rightNode
		}

		// 删除的节点有左子树，右子树为空
		if node.right == nil {
			leftNode := node.left
			node.left = nil
			b.size--
			return leftNode
		}

		// 删除的节点左右子树不为空
		// 找到比待删除节点大的最小节点，即待删除节点右子树的最小节点
		// 用这个节点顶替待删除节点的位置
		successor := b.minimum(node.right) // 找到后继
		successor.right = b.removeMin(node.right)
		// b.size++
		successor.left = node.left

		node.left = nil
		node.right = nil
		// b.size--

		return successor


		// 寻找前驱来删除左右子树不为空的写法
		//predecessor := b.maximum(node.left) // 找到前驱
		//predecessor.left = b.removeMax(node.left)
		//predecessor.right = node.right
		//
		//node.left = nil
		//node.right = nil
		//
		//return predecessor
	}

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