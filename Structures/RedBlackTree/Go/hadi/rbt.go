/*
 * +------------------------------------+
 * | Author:        hadi abbasi	        |
 * |                                    |
 * | Link:          github.com/hawwwdi  |
 * +------------------------------------+
 */

package rbt

import (
	"errors"
)

type color int8

const (
	Black color = iota
	Red
)

type Node struct {
	Value  int
	color  color
	parent *Node
	left   *Node
	right  *Node
}

func (n *Node) BlackHeight() int {
	return 0
}

type RBT struct {
	root *Node
	null *Node
}

func NewRBT() *RBT {
	rbt := &RBT{
		null: &Node{
			Value: -1,
			color: Black,
		},
	}
	return rbt
}

//Height max = 2log(n+1)
func (t *RBT) Height() int {
	return 0
}

func (t *RBT) Insert(key int) error {
	return rbtInsert(t, key)
}

func uncle(t *RBT, n *Node) *Node {
	if n.parent == t.null || n.parent.parent == t.null {
		return t.null
	}
	if n.parent.parent.left == n.parent {
		return n.parent.parent.right
	}
	return n.parent.parent.left
}

func leftRotate(t *RBT, n *Node) error {
	y := n.right
	// if y == nil {
	// 	return errors.New("right child is nil")
	// }
	n.right = y.left
	if y.left != t.null {
		y.left.parent = n
	}
	if n == t.root {
		t.root = y
	} else if n.parent.right == n {
		n.parent.right = y
	} else {
		n.parent.left = y
	}
	y.left = n
	y.parent = n.parent
	n.parent = y
	return nil
}

func rightRotate(t *RBT, n *Node) error {
	x := n.left
	// if x == nil {
	// 	return errors.New("left child is nil")
	// }
	n.left = x.right
	if x.right != t.null {
		x.right.parent = n
	}
	if n == t.root {
		t.root = x
	} else if n.parent.right == n {
		n.parent.right = x
	} else {
		n.parent.left = x
	}
	x.right = n
	x.parent = n.parent
	n.parent = x
	return nil
}

func rbtInsert(t *RBT, key int) error {
	node, err := treeInsert(t, key)
	if err != nil {
		return err
	}
	colorFixUp(t, node)
	return nil
}

func colorFixUp(t *RBT, n *Node) {
	if n == t.root {
		n.color = Black
		return
	}
	if n.parent.color == Black {
		return
	}
	nodeUncle := uncle(t, n)
	if nodeUncle.color == Red {
		nodeUncle.color = Black
		n.parent.color = Black
		n.parent.parent.color = Red
	} else {
		if isLeftChild(t, n) && isLeftChild(t, n.parent) {
			n.parent.color, n.parent.parent.color = n.parent.parent.color, n.parent.color
			rightRotate(t, n.parent.parent)
		} else if isRightChild(t, n) && isLeftChild(t, n.parent) {
			leftRotate(t, n.parent)
			n.parent.color, n.parent.parent.color = n.parent.parent.color, n.parent.color
			rightRotate(t, n.parent)
		} else if isRightChild(t, n) && isRightChild(t, n.parent) {
			n.parent.color, n.parent.parent.color = n.parent.parent.color, n.parent.color
			leftRotate(t, n.parent.parent)
		} else {
			rightRotate(t, n.parent)
			n.parent.color, n.parent.parent.color = n.parent.parent.color, n.parent.color
			leftRotate(t, n.parent)
		}
	}
	colorFixUp(t, n.parent.parent)
	return
}

func treeInsert(t *RBT, key int) (*Node, error) {
	node := &Node{
		Value: key,
		color: Red,
		right: t.null,
		left:  t.null,
	}
	if t.root == nil {
		t.root = node
		t.root.parent = t.null
		t.null.left = t.root
		t.null.right = t.root
		return node, nil
	}
	var curr, next *Node
	next = t.root
	for next != t.null {
		curr = next
		if key > curr.Value {
			next = curr.right
		} else if key < curr.Value {
			next = curr.left
		} else {
			return nil, errors.New("key already exists")
		}
	}
	if key > curr.Value {
		curr.right = node
		return node, nil
	}
	curr.left = node
	return node, nil
}

func treeDelete(t *RBT, n *Node) *Node {
	if isRoot(t, n) {
		t.root = nil
		return n
	}
	if isLeaf(t, n) {
		if isRightChild(t, n) {
			n.parent.right = t.null
			return n
		}
		n.parent.left = t.null
		return n
	}
	if hasOnlyOneChild(t, n) {
		if hasRightChild(t, n) {
			if isRightChild(t, n) {
				n.parent.right = n.right
			} else {
				n.parent.left = n.right
			}
			n.right.parent = n.parent
		} else {
			if isRightChild(t, n) {
				n.parent.right = n.left
			} else {
				n.parent.left = n.left
			}
			n.left.parent = n.parent
		}
		return n
	}
	successor := inorderSeccessor(t, n)
	n.Value = successor.Value
	//todo check color
	return treeDelete(t, successor)
}

func min(t *RBT, node *Node) *Node {
	for ; node.left != t.null; node = node.left {
	}
	return node
}

func max(t *RBT, node *Node) *Node {
	for ; node.right != t.null; node = node.right {
	}
	return node
}

func inorderSeccessor(t *RBT, node *Node) *Node {
	if node.right != nil {
		return min(t, node.right)
	}
	if node == t.root {
		return nil
	}
	curr := node.parent
	for ; curr != nil && curr.Value < node.Value; curr = curr.parent {
	}
	return curr
}

func inorderPreDecessor(t *RBT, node *Node) *Node {
	if node.left != nil {
		return max(t, node.left)
	}
	if node == t.root {
		return nil
	}
	curr := node.parent
	for ; curr != nil && curr.Value > node.Value; curr = curr.parent {
	}
	return curr
}

func isRoot(t *RBT, n *Node) bool {
	return t.root == n
}

func isRightChild(t *RBT, n *Node) bool {
	if n.parent == t.null {
		return false
	}
	return n.parent.right == n
}

func isLeftChild(t *RBT, n *Node) bool {
	if n == t.null {
		return false
	}
	return n.parent.left == n
}

func isLeaf(t *RBT, n *Node) bool {
	return n.right == t.null && n.left == t.null
}

func hasRightChild(t *RBT, n *Node) bool {
	return n.right != t.null
}

func hasLeftChild(t *RBT, n *Node) bool {
	return n.left != t.null
}

func hasOnlyOneChild(t *RBT, n *Node) bool {
	return (hasRightChild(t, n) && !hasLeftChild(t, n)) || (hasLeftChild(t, n) && !hasRightChild(t, n))
}
