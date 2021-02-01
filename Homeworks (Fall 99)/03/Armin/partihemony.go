package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

// function for scanning
func scan(index, left, right *int) {
	line, _ := reader.ReadString('\n')
	line = strings.TrimSuffix(line, "\n")
	numbers := strings.Split(line, " ")
	*index, _ = strconv.Atoi(numbers[0])
	*left, _ = strconv.Atoi(numbers[1])
	*right, _ = strconv.Atoi(numbers[2])
}

// just binary search tree implementation , main algorithm is in main func
type Node struct {
	Parent *Node
	Left   *Node
	Right  *Node
	enter  int
	out    int
	money  int
}

type Tree struct {
	Root *Node
}

func CreateTree() *Tree {
	tree := &Tree{
		Root: nil,
	}
	return tree
}

func (tree *Tree) Insert(enter, out, money int) {
	newNode := &Node{
		Parent: nil,
		Left:   nil,
		Right:  nil,
		enter:  enter,
		out:    out,
		money:  money,
	}
	if tree.Root == nil {
		tree.Root = newNode
	} else {
		tmp := tree.Root
		for {
			holder := tmp
			if money > tmp.money {
				tmp = tmp.Right
				if tmp == nil {
					holder.Right = newNode
					newNode.Parent = holder
					break
				}
			} else {
				tmp = tmp.Left
				if tmp == nil {
					holder.Left = newNode
					newNode.Parent = holder
					break
				}
			}
		}
	}
}
func (root *Node) GetMax() *Node {
	tmp := root
	for {
		holder := tmp
		tmp = tmp.Right
		if tmp == nil {
			return holder
		}
	}
}

func (root *Node) GetMin() *Node {
	tmp := root
	for {
		holder := tmp
		tmp = tmp.Left
		if tmp == nil {
			return holder
		}
	}
}

func (node *Node) GetPredecessor() *Node {
	if node.Left != nil {
		return node.Left.GetMax()
	} else {
		tmp := node
		tmp2 := tmp.Parent
		for tmp2 != nil {
			if tmp != tmp2.Left {
				break
			}
			tmp = tmp2
			tmp2 = tmp2.Parent
		}
		return tmp2
	}
}

func (node *Node) GetSuccessor() *Node {
	if node.Right != nil {
		return node.Right.GetMin()
	} else {
		tmp := node
		tmp2 := tmp.Parent
		for tmp2 != nil {
			if tmp != tmp2.Right {
				break
			}
			tmp = tmp2
			tmp2 = tmp2.Parent
		}
		return tmp2
	}
}

func (node *Node) DeleteNode(tree *Tree) bool {
	if node == nil {
		return false
	} else {
		if node.IsLeaf() { //node has no children
			if node.IsRightChild() {
				node.Parent.Right = nil
			} else {
				node.Parent.Left = nil
			}
			node.Parent = nil
			//    tree.DisplayTree()
		} else if node.Right != nil && node.Left == nil { //node has one children at right
			node.Right.Parent = node.Parent
			if node.IsRightChild() {
				node.Parent.Right = node.Right
			} else {
				node.Parent.Left = node.Right
			}
		} else if node.Left != nil && node.Right == nil { //node has one children at left
			node.Left.Parent = node.Parent
			if node.IsRightChild() {
				node.Parent.Right = node.Left
			} else {
				node.Parent.Left = node.Left
			}
		} else { //node has two children
			sucNod := node.GetSuccessor()
			holder1 := sucNod.money
			holder2 := sucNod.enter
			holder3 := sucNod.out

			sucNod.money = node.money
			sucNod.enter = node.enter
			sucNod.out = node.out

			node.money = holder1
			node.enter = holder2
			node.out = holder3

			sucNod.DeleteNode(tree)
		}
		return true
	}
}

func (node *Node) IsLeaf() bool {
	if node.Right == nil && node.Left == nil {
		return true
	}
	return false
}
func (node *Node) IsRightChild() bool {
	if node == node.Parent.Right {
		return true
	}
	return false
}

//////////////////////////////////////////////////////////end of implementation

func main() {
	reader = bufio.NewReader(os.Stdin)
	tree := CreateTree()
	var n int
	var m int
	fmt.Scan(&n)
	fmt.Scan(&m)
	var enter int
	var out int
	var money int
	for i := 0; i < n; i++ {
		scan(&enter, &out, &money)
		tree.Insert(enter, out, money)
	}
	for j := 1; j <= m; j++ {
		flag := false
		guest := tree.Root.GetMax()
		for guest != nil {
			if guest.out >= j && guest.enter <= j {
				flag = true
				break
			}
			if guest.out < j && guest != tree.Root {
				guest.DeleteNode(tree)
				guest = tree.Root.GetMax()
			} else {
				guest = guest.GetPredecessor()
			}
		}
		if !flag {
			fmt.Print("0 ")
		} else {
			fmt.Print(guest.money)
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
