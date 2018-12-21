package binary_search_tree

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

	if b.root == nil {
		b.root = NewNode(e)
		b.size++
	} else {
		b.add(b.root,e)
	}
}

// 向以node为根的二分搜索树中插入元素e，递归算法
func (b *BST)add(node *Node,e int)  {

	if node.e == e {
		return
	} else if node.left == nil && node.e > e{
		node.left = NewNode(e)
		b.size++
		return
	} else if node.right == nil && node.e < e {
		node.right = NewNode(e)
		b.size++
		return
	}

	if node.e > e {
		b.add(node.left,e)
	} else { // node.e < e
		b.add(node.right,e)
	}
}