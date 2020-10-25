/*
 * +------------------------------------+
 * | Author:        hadi abbasi	        |
 * |                                    |
 * | Link:          github.com/hawwwdi  |
 * +------------------------------------+
 */

package binarysearchtree

import (
	"fmt"
	"math"
)

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

func (n *Node) IsSubBalanced() bool {
	return isBalanced(n)
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
	if n.right == nil {
		return nil
	}
	return min(n.right)
}

func (n *Node) Predecessor() *Node {
	if n.left == nil {
		return nil
	}
	return max(n.left)
}

func (n *Node) SubDepth() int {
	return depth(n)
}

func (n *Node) inorder() {
	inorder(n)
}

func (n *Node) SubPrint() {
	//print("", n, false)
	pprint(n, 0)
}

type BST struct {
	root *Node
	size int
}

func NewBST() *BST {
	return new(BST)
}

func (b *BST) Insert(values ...int) {
	for _, v := range values {
		b.size++
		if b.root == nil {
			n := new(Node)
			n.Value = v
			b.root = n
			continue
		}
		insert(v, b.root)
	}
}

func (b *BST) Search(key int) *Node {
	return search(key, b.root)
}

func (b *BST) IsBalanced() bool {
	return isBalanced(b.root)
}

func (b *BST) Root() *Node {
	return b.root
}

func (b *BST) Max() *Node {
	return max(b.root)
}

func (b *BST) Min() *Node {
	return min(b.root)
}

func (b *BST) Depth() int {
	return depth(b.root)
}

func (b *BST) Size() int {
	return b.size
}

func (b *BST) Inorder() {
	inorder(b.root)
}

func (b *BST) Print() {
	//print("", b.root, false)
	pprint(b.root, 0)
}

func insert(value int, root *Node) {
	var curr, next *Node
	next = root
	for next != nil {
		curr = next
		if value > next.Value {
			next = next.right
		} else if value < next.Value {
			next = next.left
		} else {
			break
		}
	}
	n := new(Node)
	n.Value = value
	if value > curr.Value {
		curr.right = n
	} else if value < curr.Value {
		curr.left = n
	} else {
		return
	}
	n.parent = curr
}

/* func rinsert(value int, parent *Node) {
	if value < parent.Value {
		if parent.left != nil {
			insert(value, parent.left)
			return
		}
		n := new(Node)
		n.parent = parent
		n.Value = value
		parent.left = n
	} else if value > parent.Value {
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
*/
//todo
/*
func remove(key int, root *Node) *Node {
	node := search(key, root)
	if node == nil {
		return nil
	}
	if node.IsLeaf() {

	}
} */

func search(key int, root *Node) *Node {
	if root == nil || root.Value == key {
		return root
	}
	if key < root.Value {
		return search(key, root.left)
	} else {
		return search(key, root.right)
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

func depth(node *Node) int {
	if node == nil {
		return 0
	}
	right := depth(node.right)
	left := depth(node.left)
	if right > left {
		return right + 1
	} else {
		return left + 1
	}

}

func isBalanced(node *Node) bool {
	if node == nil {
		return true
	}
	left := depth(node.left)
	right := depth(node.right)
	if math.Abs(float64(left-right)) > 1 {
		return false
	}
	return isBalanced(node.left) && isBalanced(node.right)
}

func inorder(node *Node) {
	if node != nil {
		inorder(node.left)
		fmt.Println(node.Value)
		inorder(node.right)
	}
}

/*
func print(prefix string, n *Node, isLeft bool) {
	if n != nil {
		var s1, s2 string
		if isLeft {
			s1 = fmt.Sprint(prefix, "|-- ", n.Value)
			s2 = fmt.Sprint(prefix, "|   ")
		} else {
			s1 = fmt.Sprint(prefix, "\\-- ", n.Value)
			s2 = fmt.Sprint(prefix, "    ")
		}
		fmt.Println(s1)
		print(s2, n.left, true)
		print(s2, n.right, false)
	}
} */

func pprint(node *Node, space int) {
	if node == nil {
		return
	}
	space += 7
	pprint(node.right, space)
	fmt.Println()
	for i := 7; i < space; i++ {
		fmt.Print(" ")
	}
	fmt.Println(node.Value)
	pprint(node.left, space)
}
