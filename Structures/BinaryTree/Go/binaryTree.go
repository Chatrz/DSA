package binarytree

type Node struct {
	parent      *Node
	left, right *Node
	Value       int
}

func (n *Node) IsLeaf() bool {
	return n.left == nil && n.right == nil
}

func (n *Node) IsParent() bool {
	return n.parent == nil
}

func (n *Node) Left() *Node {
	return n.left
}

func (n *Node) Right() *Node {
	return n.right
}

func (n *Node) SubSearch(value int) *Node {
	return search(value, n)
}

func (n *Node) SubMin() *Node {
	return min(n)
}

func (n *Node) SubMax() *Node {
	return max(n)
}

func (n *Node) Seccessor() *Node {
	curr := n.right
	if curr == nil {
		return nil
	}
	return min(curr)
}

func (n *Node) Predecessor() *Node {
	curr := n.left
	if curr == nil {
		return nil
	}
	return max(curr)
}

type BT struct {
	root *Node
	size int
}

func NewBT() *BT {
	return new(BT)
}

func (b *BT) Insert(value int) {
	if b.root == nil {
		n := new(Node)
		n.Value = value
		b.root = n
		return
	}
	insert(value, b.root)
}

func (b *BT) Search(key int) *Node {
	return search(key, b.root)
}

func (b *BT) Root() *Node {
	return b.root
}

func (b *BT) Max() *Node {
	return max(b.root)
}

func (b *BT) Min() *Node {
	return min(b.root)
}

func insert(value int, parent *Node) {
	if value < parent.Value {
		if parent.left != nil {
			insert(value, parent.left)
			return
		}
		n := new(Node)
		n.parent = parent
		n.Value = value
		parent.left = n
	} else {
		if parent.right != nil {
			insert(value, parent.right)
			return
		}
		n := new(Node)
		n.parent = parent
		n.Value = value
		parent.right = n
	}
}

func search(key int, node *Node) *Node {
	if node == nil || node.Value == key {
		return node
	}
	if key < node.Value {
		return search(key, node.left)
	} else {
		return search(key, node.right)
	}
}

func min(node *Node) *Node {
	for ; node.left != nil; node = node.left {
	}
	return node
}

func max(node *Node) *Node {
	for ; node.right != nil; node = node.right {
	}
	return node
}
