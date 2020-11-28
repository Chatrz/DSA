/*
 * +------------------------------------+
 * | Author:        hadi abbasi	        |
 * |                                    |
 * | Link:          github.com/hawwwdi  |
 * +------------------------------------+
 */

package rbt

import "errors"

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

func (n *Node) Uncle() *Node {
	return uncle(n)
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
func (rbt *RBT) Height() int {
	return 0
}

func uncle(n *Node) *Node {
	if n.parent == nil || n.parent.parent == nil {
		return nil
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
	if y.left != nil {
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
	if x.right != nil {
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

func treeInsert(t *RBT, key int) error {
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
		t.root.color = Black
		return nil
	}
	var curr, next *Node
	next = t.root
	for next != nil {
		curr = next
		if key > curr.Value {
			next = curr.right
		} else if key < curr.Value {
			next = curr.left
		} else {
			return errors.New("key already exists")
		}
	}
	if key > curr.Value {
		curr.right = node
		return nil
	}
	curr.left = node
	return nil
}
