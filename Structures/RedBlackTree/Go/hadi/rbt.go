/*
 * +------------------------------------+
 * | Author:        hadi abbasi	        |
 * |                                    |
 * | Link:          github.com/hawwwdi  |
 * +------------------------------------+
 */

package rbt

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
}

func NewRBT() *RBT {
	rbt := &RBT{}
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
